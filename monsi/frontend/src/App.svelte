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
    <h3>Monsi</h3>
    <div class="row">
        <div class="column">
            <div id="DIDlist"></div>
                {#each DIDs as did, i}
                    <button on:click={() => getListOfVCs({i})}>{did.did} </button><br/>
                {/each}
            <div id="result">---</div>
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