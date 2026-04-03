package main

type Team struct{
 TeamId string `json:"tid"`
 TeamName string `json:"tname"`
 TeamPrice int `json:"price"`
 Hero *Hero `json:"hero"`
}

type Hero struct{
 FullName string `json:"fullname"`
 Website string `json:"website"`
}

// =--=- FAke DB ==--=-//

var team[] Team;


//=--=-= Middleware, helper -=-==-//
func (t *Team) IsEmpty() bool{
	return t.TeamId =="" && t.TeamName=="";
}


func main(){

}