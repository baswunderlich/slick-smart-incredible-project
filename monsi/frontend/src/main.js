import './style.css';
import './app.css';

import {AddDID} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
    <h1>Here your DIDs should be visible</h1>
    <h1>Here others DIDs and publicKeys should be visible</h2>
    <div id="result">---</div>
    <button class="btn" onclick="addDID()">Add DID</button>
`;

let resultElement = document.getElementById("result");

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
