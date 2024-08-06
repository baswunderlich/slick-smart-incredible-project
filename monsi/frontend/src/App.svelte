<script>
    import { GetListOfDIDs } from "../wailsjs/go/main/App.js";
    import { GetListOfVCs } from "../wailsjs/go/main/App.js";
    import { AddDID } from "../wailsjs/go/main/App.js";
	import { setContext } from 'svelte';
	import { writable } from 'svelte/store';
	import { getContext } from 'svelte';

    let DIDs = [];
    let VCs = [];               //context: "vcs"
    let selectedDID = "";       //context: "selectedDID"

    let getListOfDIDs = function() {
        let newDIDs = [];
        console.log("get list of DIDS");
        try{
            GetListOfDIDs()
                .then((result) => {
                    DIDs = result;
                })
                .catch((err) => {
                    console.error(err);
                });
        } catch (err) {
            console.error(err);
        }
    }

    // Setup the addDID function
    let addDID = function () {
        let resultElement = document.getElementById("result");
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

    let getListOfVCs = function(i) {
        let index = i.i;
        console.log(index);
        let did = DIDs[index];
        console.log("get list of VCs of DID " + did);
        try{
            GetListOfVCs(did)
                .then((result) => {
                    VCs = result;
                })
                .catch((err) => {
                    console.error(err);
                });
        } catch (err) {
            console.error(err);
        }
    }

    function init() {
        getListOfDIDs();
    }

    init();
</script>

<main>
    <h3>Monsi</h3>
    <div class="row">
        <div class="column">
            <div id="DIDlist"></div>
                {#each DIDs as did, i}
                    <button on:click={() => getListOfVCs({i})}>{did} <br/></button>
                {/each}
            <div id="result">---</div>
            <button class="btn" on:click="{addDID}">Add DID</button>
            DID: <input id="DIDField"/>     
        </div>
        <div class="column">  
            {#each VCs as vc}
                {vc} <br/>
            {/each} 
        </div> 
    </div>
</main>

<style>
    .row {
    display: flex;
    }

    .column {
    flex: 50%;
    }
</style>