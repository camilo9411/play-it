var firstPlayer = true;

document.addEventListener("DOMContentLoaded", function(){
    

});

function clickHandler(column){
    // alert("your pressed this: "+column)

    fetch("yellowRed.html/play?c=" + column)
    .then(response => response.json())
    .then(data => {

        if(data.result == true){
            location.reload();
        }
        if(data.row_choice == -1){
            alert("You can not add in this column!")
        }else{
            //we have a row choice
            //changing visual for players
            if(firstPlayer){
                document.getElementById("player_turn").innerHTML = "Player 2's turn";
                document.getElementById("player_turn_div").style.backgroundColor="yellow";
            }else{
                document.getElementById("player_turn").innerHTML = "Player 1's turn";
                document.getElementById("player_turn_div").style.backgroundColor="Red";
            }

            //showing the turn
            result = "cellr"+data.row_choice+"c"+data.column_choice;
            document.getElementById(result).style.backgroundColor = firstPlayer ? "red" : "yellow";
            
            //here checking if we have winner
            if (data.winner != 0){
                alert(data.msg);
            }

            firstPlayer = !firstPlayer
        }
    })
}


function resetGame(){

    location.reload("yellowRed.html");
}