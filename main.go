package main

// import (
// 	"fmt"
// 	// "sort"
// "math/rand"

// )

// // 1-EXERCISES ..exam..

// func RandomN(n int) {

//     ch1 := make(chan int)
//     ch2 := make(chan int)

//     go func() {
//         defer close(ch1)
//         for i := 0; i < n; i++ {
//             ch1 <- rand.Intn(21) + 10
//         }
//     }()

//     go func() {
//         defer close(ch2)
//         for num := range ch1 {
// 		fmt.Println("the number is",num)
//             ch2 <- num * num

//         }
//     }()

//     for val := range ch2 {
//       fmt.Println(val)
//     }
// }

// func main() {
//     RandomN(10)
// }

// 2-EXERCISES ..exam..
// SELECT e.name, d.name, SUM(s.amount),count(s.amount) as How_many_times
// FROM employees e
// JOIN salaries s ON e.id = s.employee_id
// JOIN departments d ON e.department_id = d.id
// GROUP BY e.name, d.name, amount;

// 3-EXERCISES ..exam..
// SELECT d.name,e.salary,e.name
// FROM employee e
// JOIN department d ON e.departmentId = d.id where salary > 70000;

// 4-EXERCISES ..exam..
// func main()  {
//     var str2 = []string{"dog","racecar","car"}
//     var str = []string{"flower","flow","flight"}
//     fmt.Println(prefixwords(str))
//     fmt.Println(prefixwords(str2))
// }

// func prefixwords(strs []string) string {
//     var long string = ""
//     var end = false
    
//     if len(strs) > 0 {
//         sort.Strings(strs)
//         sum := string(strs[0])
//         last := string(strs[len(strs)-1])
        
//         for i := 0; i < len(sum); i++ {
//             if !end && string(last[i]) == string(sum[i]) {
//                 long += string(last[i])
//             } else {
//                 end = true
//             }
//         }
//     }
//     return long
// }