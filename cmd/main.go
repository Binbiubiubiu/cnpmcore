package main

import "time"

func main() {
	// var repo = &repository.ChangeRepository{}
	// // var newB = &entity.Change{
	// // 	ChangeId:   util.CreateObjectId(),
	// // 	Type:       "Category",
	// // 	TargetName: "Parent",
	// // 	Data: map[string]any{
	// // 		"name": "name323",
	// // 		"age":  122,
	// // 	},
	// // }
	// // newB.ID = 48
	// list, err := repo.Query(0, 10)
	// if err != nil {
	// 	return
	// }
	// bs, _ := json.Marshal(list)
	// var out bytes.Buffer
	// json.Indent(&out, bs, "", "\t")
	// fmt.Printf("student=%v\n", out.String())
	println(time.Now().Format("021514"))
}
