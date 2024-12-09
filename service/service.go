package service

import (
  "fmt"
  "net/http"
  "encoding/json"
  "html/template"

)
type Service struct{
   Name string `json:"name"`
   Id int  `json:"id"`
   Owner string `json:"owner"`
   Responsible string `json:"responsible"`
   Platform string `json:"planform"`
   Purpose string `json:"purpose"`
}
var srvList []Service


func Welcome(w http.ResponseWriter, r *http.Request){
  tf,err:=template.ParseFiles("templates/service.html")
  if err!=nil{
    fmt.Fprint(w,"Error:",err)
    return
  }
  if tf==nil{
    fmt.Fprint(w,"error while creating html")
    return
  }
  err=tf.Execute(w,srvList)
  if err!=nil{
    fmt.Fprint(w,"Error:",err)
    return
  }
}

func Create(w http.ResponseWriter, r *http.Request){
  var srv Service
  err:=json.NewDecoder(r.Body).Decode(&srv)
  srvList=append(srvList,srv)
  if err!=nil{
    fmt.Fprint(w,"Unable to write",err)
    return
  }
  fmt.Fprint(w,"Service Created :", srv.Name)
}

func List(w http.ResponseWriter, r *http.Request){

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(srvList)


}
