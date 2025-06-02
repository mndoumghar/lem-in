package ants

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ants struct {
	AntNum         int
	RoomsWithCords map[string][]string
	Start          string
	End            string
	Links          []string
	Err            error
}



func ReadFile(filename string) Ants {
	var ant Ants

	// 1. قراءة الملف كاملاً كأسلسلة بايت وطباعة المحتوى
	data, err := os.ReadFile(filename)
	if err != nil {
		ant.Err = err
		return ant
	}
	fmt.Println(string(data) + "\n")

	// 2. تقسيم النص إلى أسطر، ثم إزالة المسافات الزائدة من كل سطر
	rawLines := strings.Split(string(data), "\n")
	lines := make([]string, 0, len(rawLines))
	for _, l := range rawLines {
		lines = append(lines, strings.TrimSpace(l))
	}

	// 3. تحويل السطر الأول (غير فارغ/تعليق) إلى عدد النمل
	if len(lines) == 0 {
		ant.Err = fmt.Errorf("invalid data format, empty file")
		return ant
	}
	// نفرض أن السطر الأول هو رقم النمل مباشرة
	ant.AntNum, err = strconv.Atoi(lines[0])
	if err != nil {
		ant.Err = err
		return ant
	}
	if ant.AntNum <= 0 {
		ant.Err = fmt.Errorf("invalid data format, no ant in the start room")
		return ant
	}

	// 4. تهيئة الخرائط والمصفوفات
	ant.RoomsWithCords = make(map[string][]string)
	ant.Links = []string{}
	roomNames := make(map[string]bool)    // للتحقّق من تكرار أسماء الغرف
	coordinates := make(map[string]bool)  // للتحقّق من تكرار الإحداثيات

	// 5. المرور على باقي الأسطر (بدءًا من index=1)
	//    والتعامل مع كل حالة: ##start، ##end، غرف، روابط، تعليقات، أسطر فارغة
	i := 1
	for i < len(lines) {
		line := lines[i]

		// 5.1 إذا كان السطر خالي أو تعليق (غير ##start/##end)، نتجاهله
		if line == "" || (strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##")) {
			i++
			continue
		}

		// 5.2 التعامل مع ##start أو ##end
		if line == "##start" || line == "##end" {
			isStart := (line == "##start")
			// نتأكد أن هناك سطر تالي لتعريف الغرفة
			if i+1 >= len(lines) {
				ant.Err = fmt.Errorf("invalid data format, missing %s room definition", map[bool]string{true: "start", false: "end"}[isStart])
				return ant
			}

			next := strings.Fields(lines[i+1])
			if len(next) != 3 {
				ant.Err = fmt.Errorf("invalid data format, invalid %s room", map[bool]string{true: "start", false: "end"}[isStart])
				return ant
			}
			name := next[0]
			xStr, yStr := next[1], next[2]
			X, errX := strconv.Atoi(xStr)
			Y, errY := strconv.Atoi(yStr)
			if errX != nil || errY != nil {
				ant.Err = fmt.Errorf("invalid data format, invalid coordinates")
				return ant
			}
			
			coordKey := fmt.Sprintf("%d,%d", X, Y)
			// التحقق من اسم الغرفة والإحداثيات
			if strings.HasPrefix(name, "#") || strings.HasPrefix(name, "L") || roomNames[name] || coordinates[coordKey] {
				ant.Err = fmt.Errorf("invalid data format, duplicate or invalid room/%s", name)
				return ant
			}
			// الحفظ
			roomNames[name] = true
			coordinates[coordKey] = true
			ant.RoomsWithCords[name] = []string{strconv.Itoa(X), strconv.Itoa(Y)}
			if isStart {
				ant.Start = name
			} else {
				ant.End = name
			}
			// نزيد i مرتين، لأننا عالجنا السطر الحالي وسطر التعريف التالي
			i += 2
			continue
		}

		// 5.3 تقسيم السطر إلى أجزاء لمعرفة إن كان غرفة عادية أو رابط
		parts := strings.Fields(line)

		// 5.3.1 إذا كان عدد الأجزاء 3 => وصف غرفة عادية
		if len(parts) == 3 {
			name := parts[0]
			xStr, yStr := parts[1], parts[2]
			X, errX := strconv.Atoi(xStr)
			Y, errY := strconv.Atoi(yStr)
			if errX != nil || errY != nil {
				ant.Err = fmt.Errorf("invalid data format, invalid coordinates")
				return ant
			}
			coordKey := fmt.Sprintf("%d,%d", X, Y)
			if strings.HasPrefix(name, "#") || strings.HasPrefix(name, "L") || roomNames[name] || coordinates[coordKey] {
				ant.Err = fmt.Errorf("invalid data format, duplicate or invalid room/%s", name)
				return ant
			}
			roomNames[name] = true
			coordinates[coordKey] = true
			ant.RoomsWithCords[name] = []string{strconv.Itoa(X), strconv.Itoa(Y)}
			i++
			continue
		}

		// 5.3.2 إذا كان السطر يحتوي على "-" (دون أن يبدأ بـ "#") => نعتبره رابطًا
		if strings.Contains(parts[0], "-") && !strings.HasPrefix(parts[0], "#") && len(parts) == 1 {
			ant.Links = append(ant.Links, parts[0])
			i++
			continue
		}

		// 5.3.3 أي حالة أخرى غير متوقعة (تنسيق غير معروف) => نتجاهلها أو نعيد خطأ
		// هنا سنعدها خطأ في التنسيق
		ant.Err = fmt.Errorf("invalid data format, unrecognized line: %q", line)
		return ant
	}

	// 6. التحقق من وجود غرفتي البداية والنهاية وعدم تطابقهما
	if ant.Start == "" || ant.End == "" || ant.Start == ant.End {
		ant.Err = fmt.Errorf("invalid data format, missing start or end room")
		return ant
	}

	// 7. التحقق من عدم وجود روابط بين ##start و ##end في نفس الكتلة
	if !isBetween(filename) {
		ant.Err = fmt.Errorf("invalid data format, paths shouldn't be between ##start and ##end rooms")
		return ant
	}

	return ant
}

// دالة مساعدة للتحقق من الروابط الواردة بين كتلة ##start و ##end
func isBetween(filename string) bool {
	fh, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer fh.Close()

	inStartBlock := false
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "##start" {
			inStartBlock = true
			continue
		}
		if line == "##end" {
			inStartBlock = false
			continue
		}
		// إذا كنّا داخل كتلة ##start ولم نصل بعد إلى ##end ورأينا رابطاً => خطأ
		if inStartBlock && strings.Contains(line, "-") {
			return false
		}
	}
	return scanner.Err() == nil
}
