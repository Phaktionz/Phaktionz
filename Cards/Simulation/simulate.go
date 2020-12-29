package main

import (
	"fmt"
)

type Players struct{
	T3 int
	T2 int
	T1 int
	Reg int
	C int;
	W int
	R int
}
func check (p Players)bool{
	sum := 0
	sum += p.T3 + p.T2 + p.T1 + p.Reg + p.C + p.W+p.R
	if sum == 50 {
		return true
	} else {
		return false
	}
}
func main(){
		Zane := Players{
			T3:  0,
			T2:  16,
			T1:  20,
			Reg: 10,
			C:   3,
			W:   0,
			R:   1,
		}
		
		Lulo := Players{
			T3:  4,
			T2:  12,
			T1:  20,
			Reg: 4,
			C:   7,
			W:   2,
			R:   1,
		}
		//Check if decks qualify and quantities are right
		if check(Zane) && check(Lulo) == true{
			//Run
			var Z1, Z2,Z3,ZReg, ZC, ZW, ZR int
			Z3 = Zane.T3
			Z2 = Zane.T2
			Z1 = Zane.T1
			ZReg = Zane.Reg
			ZC = Zane.C
			ZW = Zane.W
			ZR = Zane.R

			var Z[6] int

			for i:=0; i<6; i++ {
				if Z[i]  Z3 {

				}
			}







		}else {
			fmt.Println("Not sufficient Cards to Simulate")
		}

	}
