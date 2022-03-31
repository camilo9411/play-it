var count = 0
function CheckAnswer(){
    let answer = "'" + document.querySelector("input").value + "'";
     
    for (let node of document.querySelectorAll("td")){
        if (node.className.indexOf(answer) >= 0){
            node.style = "color: black;";
        }
    }
    count++
    
    if (count > 5){
        alert("Congrats !")
    }
    document.querySelector("input").value = "";

    
}