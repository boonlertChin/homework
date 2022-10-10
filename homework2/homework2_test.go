package homework2

import ("testing")

func TestIsTailSum(t *testing.T) { 
	
	/*
	testInput  =  int[] {1, 2, 3, 4, 5} 
	ถ้ากำหนดผลรวม k = 12 จะได้ค่า n = 3 เพราะ 3 ลำดับสุดท้ายรวมกันได้ 12 
	*/
	testInput := []int{1, 2, 3, 4, 5} 
	k:= 12
	n := 3
	got := IsTailSum(testInput,k)
    if got != n {
        t.Errorf("IsTailSum() = %d, want %d", got, n)
    }
	
	/*
	testInput  = int[] {1, 2, 3, 4, 5} 
	ถ้ากำหนดผลรวม k = 15 จะได้ค่า n = 5 เพราะ 5 ลำดับสุดท้ายรวมกันได้ 15
	*/
	testInput = []int{1, 2, 3, 4, 5} 
	k = 15
	n = 5
    got = IsTailSum(testInput,k)
    if got != n {
        t.Errorf("IsTailSum() = %d, want %d", got, n)
    }

	/*
	testInput  = int[] {1, 2, 3, 4, 5} 
	ถ้ากำหนดผลรวม k = 8 จะได้ค่า 
	n = 0 
	เพราะไม่มีลำดับใด ๆ รวมกันได้ 8
	*/
	testInput = []int{1, 2, 3, 4, 5} 
	k = 8
	n = 0
    got = IsTailSum(testInput,k)
    if got != n {
        t.Errorf("IsTailSum() = %d, want %d", got, n)
    }

	/*
	testInput  = int[] {1, -1, 1, -1, 1} 
	ถ้ากำหนดผลรวม k = 0 จะได้ค่า 
	n = 2 
	เพราะ 2 ลำดับสุดท้ายและ 4 ลำดับสุดท้ายรวมกันได้ 0 ให้เลือกลำดับที่น้อยกว่า
	*/
	testInput = []int{1, -1, 1, -1, 1}  
	k = 0
	n = 2
    got = IsTailSum(testInput,k)
    if got != n {
        t.Errorf("IsTailSum() = %d, want %d", got, n)
    }

	/*
	testInput  = int[] {} 
	กำหนดผลรวม K = อะไรก็ได้ค่า 
	n = 0 
	เพราะไม่มีลำดับใด ๆ รวมกันได้เท่ากับ K
	*/
	testInput = []int{}  
	k = 50 //อะไรก็ได้ค่า 
	n = 0
    got = IsTailSum(testInput,k)
    if got != n {
        t.Errorf("IsTailSum() = %d, want %d", got, n)
    }

	/*
	testInput  = int[] {3, 0, -2, 4, -6}
	ถ้ากำหนดผลรวม k = -4 จะได้ค่า 
	n = 3 
	เพราะ 3 ลำดับสุดท้ายรวมกันได้ -4
	*/
	testInput = []int{3, 0, -2, 4, -6}
	k = -4 //อะไรก็ได้ค่า 
	n = 3
    got = IsTailSum(testInput,k)
    if got != n {
        t.Errorf("IsTailSum() = %d, want %d", got, n)
    }
}