import React from 'react'
import './Pieces.css'
function Piece(props){

    let styles = {
        backgroundColor: props.color === "Red" ? "#D90368" : "#FED766"
    }

    return (
    <div className = "piece" style = {styles}>

    </div>
    );
}

export function RedPiece(){
    return <Piece color = "Red"/>
}

export function YellowPiece(){
    return <Piece color = "Yellow"/>
}