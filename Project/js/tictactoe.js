function createGridTicTacToe(){
    
    var grid = document.getElementById("tictactoe")
   
    let str = "";
    var value = 0;
    for(j=0;j<3;j++){

        str += "<div class= \"row d-flex justify-content-center\"> ";
        for(i=0;i<3;i++){

            str +=  "<button class=\"col-3 btn p-5 bg-white m-2 text-white game \"  id =\"position_" + value + "\"  value =\"" + value + "\" width=\"200px\">E</button>";
            value++;
        }
        str += "</div>"
    }

    grid.innerHTML = str

}


createGridTicTacToe();

var counter = 0;
var elements = document.getElementsByClassName("game");
var len = elements.length;

const player1 = []
const player2 = []


var singlePlayer = document.getElementById("cbSinglePlayer");

singlePlayer.addEventListener('click', (e) =>{

    resetGame();

});


for (var i = 0; i<len; i++) {
    
    elements[i].addEventListener('click', (e) => {
        
        let element = e.target;

        if(element.innerHTML == "X" || element.innerHTML == "O"){

            
        }else{
            if(counter%2 == 0 ){
                element.classList.remove('bg-white');  
                element.classList.add('bg-success');    
                element.innerHTML = "X";
                player1.push(element.value);
                if(checkWinner(player1)){
                    alert("Player 1 wins!")
                }
                
            }else{
                element.classList.remove('bg-white');  
                element.classList.add('bg-danger');
                element.innerHTML = "O";
                player2.push(element.value);
                if(checkWinner(player2)){
                    (alert("Player 2 wins!"))
                }
            }

            counter++;
            
            if(singlePlayer.checked == true && counter%2 !=0){
                playWithTheComputer(player1, player2);
            }
        }

    });
};



var winnerArray = [['0','4', '8'],['0', '1', '2'],['0','3','6'],['2','4','6'],['2','5','8'],['3','4','5'],['6','7','8'],['1','4','7']];

function checkWinner(currentPlayer){

    var row = 0;

    currentPlayer.sort();

    if(currentPlayer.length >= 3){

        for(var i=0; i < winnerArray.length; i++){
            

            for(var j=0; j < currentPlayer.length; j++){
                
                if((winnerArray[i].includes(currentPlayer[j]))){

                    row++;         
                }
            }
            if(row == 3 ){
                
                return true;
            }else{

                row = 0;
            }
        }

    }
    if(currentPlayer.length == 5){

        alert("No one wins!")
    }

    return false

}


function playWithTheComputer(currentPlayer, computerPlayer){
    
    currentPlayer.sort();
    computerPlayer.sort();
    var count = 0

    //first Move:
    if(computerPlayer.length == 0){

        switch(currentPlayer[0]){

            case '4':
                document.getElementById("position_0").click();
                break;
            default:
                document.getElementById("position_4").click();
        }
    }else{

        //Try to Win
        for(var i=0; i < winnerArray.length; i++){
                
            for(var j = 0; j < winnerArray[i].length; j++){
                
                if((computerPlayer.includes(winnerArray[i][j]))){

                    count++;

                }else{

                    positionEmpty = winnerArray[i][j];
                    
                }
                
            }

            if(!currentPlayer.includes(positionEmpty) && count >= 2){
                
                flag = true;
            }

            if(flag == true){

                let str = "position_"+ positionEmpty;
                document.getElementById(str).click();
                count =0;
                flag =  false;
                return false
            }
            
            flag =  false;
            count =0;
        }

        //Block Rival
        if(currentPlayer.length >= 2 && currentPlayer.length < 5 ){

            var flag = false
            var positionEmpty
            for(var i=0; i < winnerArray.length; i++){
                
                for(var j = 0; j < winnerArray[i].length; j++){
                    
                    if((currentPlayer.includes(winnerArray[i][j]))){

                        count++;

                    }else{

                        positionEmpty = winnerArray[i][j];
                        
                    }
                    
                }

                if(!computerPlayer.includes(positionEmpty) && count >= 2){
                    
                    flag = true;
                }

                if(flag == true){

                    let str = "position_"+ positionEmpty;
                    document.getElementById(str).click();
                    count =0;
                    flag =  false;
                    return false
                }
                
                flag =  false;
                count =0;
            }


            //Random choose
            while(true){

                let i = (Math.floor(Math.random() * 8) + 1).toString();
                console.log(i)
                if(!computerPlayer.includes(i) && !currentPlayer.includes(i)){
                    let str = "position_"+ i;
                    document.getElementById(str).click();
                    count =0;
                    flag =  false;
                    return false
                }
            }
        }

    }



}

function resetGame(){

    for (var i = 0; i<len; i++) {

        elements[i].classList.remove('bg-success');
        elements[i].classList.remove('bg-danger');  
        elements[i].classList.add('bg-white');
        elements[i].innerHTML = "E";

        while(player1.length > 0) {
            player1.pop();
        }
        while(player2.length > 0) {
            player2.pop();
        }
    }
    count=0;
    counter = 0;

}