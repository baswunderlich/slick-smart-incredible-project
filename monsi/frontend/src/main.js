import './style.css';
import './app.css';

import {AddDID} from '../wailsjs/go/main/App';
import { GetListOfDIDs } from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = 
    `<div id="DIDlist"></div>
    <h1>Here others DIDs and publicKeys should be visible</h2>
    <div id="result">---</div>
    <button class="btn" onclick="addDID()">Add DID</button>
    //Here a dialog could be added that opens when the "Add DID"-button is clicked.
`;

let resultElement = document.getElementById("result");
let listElement = document.getElementById("DIDlist");

window.getListOfDIDs = function() {
    console.log("get lsit of DIDS");
    try{
        GetListOfDIDs()
            .then((result) => {
                listElement.innerHTML = result
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
}

// Setup the addDID function
window.addDID = function () {
    try {
        AddDID()
            .then((result) => {
                // Update result with data back from App.Greet()
                resultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};


window.getListOfDIDs();