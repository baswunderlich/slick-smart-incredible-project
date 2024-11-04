<script>
    import { GetListOfDIDs } from "../wailsjs/go/main/App.js";
    import { GetListOfVCs } from "../wailsjs/go/main/App.js";
	import { setContext } from 'svelte';
	import { writable } from 'svelte/store';
	import { getContext } from 'svelte';

    let DIDs = [];
    let VCs = [];               //context: "vcs"
    let selectedDID = "";       //context: "selectedDID"

    let getListOfDIDs = function() {
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

    let getListOfVCs = function(i) {
        let index = i.i;
        console.log(index);
        let did = DIDs[index];
        console.log("get list of VCs of DID " + did);
        try{
            GetListOfVCs(did.did)
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
    <h3>Monsi was started...</h3>
    <div class="tutorial">    
        On the left you can see the available DIDs. On the right you can see the avaiable VCs matching the selected DID and until they are valid.
    </div>
    <div class="row">
        <div class="column">
            <div class="DIDlist" id="DIDlist">
                {#each DIDs as did, i}
                    <button class="button" on:click={() => getListOfVCs({i})}>{did.did} </button><br/>
                {/each}
            </div>
        </div>
        <div class="column">  
            <div class="VClist">
                {#each VCs as vc}
                    <div class="vc">{vc}</div> <br/>
                {/each} 
            </div>
        </div> 
    </div>
</main>

<style>
    .row {
    display: flex;
    padding-top: 5%;
    }

    .column {
    flex: 50%;
    }

    .VClist {
        text-align: left;
    }

    .button {
    background-color: #1d005f;
    border: none;
    color: white;
    padding: 16px 32px;
    text-align: center;
    font-size: 16px;
    margin: 4px 2px;
    opacity: 0.6;
    transition: 0.3s;
    display: inline-block;
    text-decoration: none;
    cursor: pointer;
    width: 80%;
    padding: 5%;
    border-radius: 1%;
    }
    .button:hover {opacity: 1}

    .vc {
        padding-top: 0.2em;
    }

    .tutorial {
        text-align: left;
    }
</style>