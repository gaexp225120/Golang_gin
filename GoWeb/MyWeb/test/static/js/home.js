window.onload = function() {
    
 
};

function go(){
  window.location.href = 'test.html';
}

function getuser(){


Book_Id = document.getElementById("Book_Id").value 
Book_Name = document.getElementById("Book_Name").value 

axios.post("http://127.0.0.1:8080/postbook", {
  Book_Id: Book_Id,
  Book_Name: Book_Name  ,
  responseType: "json",
})  
.then(function (response) {
  // handle success
  console.log(response);
})
.catch(function (error) {
  // handle error
  console.log(error);
})
.then(function () {
  // always executed
});
}