package homework2

func IsTailSum(testInput []int, k int) int{
	//k is Traget
	n:= 1 
	valueSummary:=0
	var i int;
	for i =len(testInput)-1 ; i>=0 ; i-- {
		valueSummary += testInput[i]
		if valueSummary == k{ //Check Summary with target
			return n //ให้เลือกลำดับที่น้อยกว่า
			break;
		}
		n++
	}
	return 0 //หากไม่พบจำนวน (n) ที่ได้ผลรวมเท่ากับตัวแปร (k) ตามที่กำหนดให้ Return 0
}

func Hello() string{
	return "Hello, World!"
}
