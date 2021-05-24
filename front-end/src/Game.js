import React, { useState } from 'react'
import './styles.css'
import axios from 'axios'
import { bounce } from 'react-animations'


/**
 * Turn state is as follows
 * 0 is for YELLOW
 * 1 is for RED
 */

var whiteStyle = {
    backgroundColor:"white"
}
var redStyle = {
    backgroundColor:"red"
}

var yellowStyle = {
    backgroundColor:"yellow"
}
export default function(){
    const [turn, setTurn] = useState(true);
    const [style, setStyle] = useState(redStyle);
    let changeTurnFunction = () => {

        if(style === redStyle) return setStyle(yellowStyle);
        else setStyle(redStyle);
        setTurn(!turn)
    }
    return (
        <div className = 'app'>

            <Board 
            changeTurn = {changeTurnFunction}
            currentTurn = {turn}
            style = {style}
            />
            

        </div>
    );
}
function CurrentTurn(props){
    return (
        <div className = "current-turn-div">
            <h1 style = {{color:props.style.backgroundColor}}>
                Current Turn
            </h1>
            <div className = "current-turn" style = {props.style}>

            </div>
        </div>
    )
}

function Board(props){
    let emptyColumns = new Array(7);
    for(let i = 0 ; i<7 ; i++){
        emptyColumns[i] = <EmptyColumn 
        style = {props.style}
        column = {i}
        currentTurn = {props.currentTurn}
        changeTurn = {props.changeTurn}
        />
    }

    return (
        <div className = "board">
            {emptyColumns}
        </div>
    );
}

function EmptyColumn(props){
    let emptySpaces = new Array(6);
    const [styleList, setStyleList] = useState({
        list: new Array(6).fill(whiteStyle)
    })
    let addPiece = (row/*, columnClicked*/) => {
    
        
        let pieceFound = false;
        let indexOfPiece = 0;

        //looks at the current emptySpace and checks the index that needs to be edited
        for(let i = 1; i < 6 && !pieceFound; i++){
            if(styleList.list[i] == whiteStyle && styleList.list[i-1] != whiteStyle){
                indexOfPiece = i;
                pieceFound = true;
            }
        }
        
        //this function should be replaced by the back end as follows
        // let webData = {
        //     user: turn ? "1" : "0",
        //     columnClicked: columnClicked
        // }
        // axios.patch("link", webData)
        // .then((res) => {
        //     let data = res.data;
        //     if(!data.gameWon)indexOfPiece = data.changedRow;
        // })

        let tempList = styleList.list;
        tempList[indexOfPiece] = props.style;

        setStyleList({
            list: tempList
        })
    }
    for(let i = 0 ; i<6 ; i++){
        emptySpaces[i] = <EmptySpace 
        style = {styleList.list[i]}
        column = {props.column} 
        row = {i}
        currentTurn = {props.currentTurn}

        onClick = {() => {
            props.changeTurn();
            addPiece(i/*, column*/);
        }}
        />
    }
    /* Reverse is done so that alignment of x and y is corrected */
    return (
    <div className = "emptyColumn">
            {emptySpaces.reverse()}
    </div>
    );
}

function EmptySpace(props){
    return (
        <div 
        className = "emptySpaces"
        style = {props.style}
        onClick = {props.onClick}
        >
        </div>
    );
}
