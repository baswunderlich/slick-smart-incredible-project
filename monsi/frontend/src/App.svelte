<script>
    import { GetListOfDIDs, RefreshVCs, RemoveVC, StoreVC } from "../wailsjs/go/main/App.js";
    import { GetListOfVCs } from "../wailsjs/go/main/App.js";

    let DIDs = [];
    let VCs = [];
    var currentDID = 0;

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

    let storeVC = function(vcName, vcContent){
        try{
            StoreVC(vcName, vcContent)
                .then(() => {
                    //Refreshing the VC list on changes  
                    getListOfVCs(currentDID)
                })
                .catch((err) => {
                    console.error(err);
                });
        }catch (err) {
            console.error(err);
        }
    }

    let getListOfVCs = function(i) {
        //This is a bit ugly and is needed beaucse of svelte 
        //not allowing to use i.i as parameter
        let index = i.i;
        if (index == undefined){
            index = i
        }

        currentDID = index;
        let did = DIDs[index];
        console.log("get list of VCs of DID " + did.did);
        try{
            GetListOfVCs(did.did)
                .then((result) => {
                    VCs = result;
                })
        } catch (err) {
            console.error(err);
        }
    }

    function init() {
        getListOfDIDs();
    }

    function dropHandler(ev) {
        console.log("File(s) dropped");

        // Prevent default behavior (Prevent file from being opened)
        ev.preventDefault();

        if (ev.dataTransfer.items) {
            // Use DataTransferItemList interface to access the file(s)
            [...ev.dataTransfer.items].forEach((item, i) => {
                // If dropped items aren't files, reject them
                if (item.kind === "file") {
                    const file = item.getAsFile();
                    file.text().then((text) => 
                    { storeVC(file.name, text) });
                }
            });
        } else {
            // Use DataTransfer interface to access the file(s)
            [...ev.dataTransfer.files].forEach((file, i) => {
                file.text().then((text) => 
                { storeVC(file.name, text) });})
        }
        
        //Refreshing the VC list on changes  
        getListOfVCs(currentDID)
    }

    function dragOverHandler(ev) {
        console.log("File(s) in drop zone");

        // Prevent default behavior (Prevent file from being opened)
        ev.preventDefault();
    }

    function removeVC(vc) {
        try{
            RemoveVC(vc)
                .then(() => {
                    //Refreshing the VC list on changes  
                    getListOfVCs(currentDID)
                })
                .catch((err) => {
                    console.error(err);
                });
        } catch (err) {
            console.error(err);
        }
    }

    function refreshVCs(){
        try{
            RefreshVCs()
                .then(() => {
                    //Refreshing the VC list on changes  
                    getListOfVCs(currentDID)
                })
                .catch((err) => {
                    console.error(err);
                });
        } catch (err) {
            console.error(err);
        }
    }

    init();
</script>

<main 
id="drop_zone"
on:drop="{(event) => dropHandler(event)}"
on:dragover="{(event) => dragOverHandler(event)}">
    <h3>Monsi was started succesfully</h3>
    <div class="tutorial">    
        On the left you can see the available DIDs. On the right you can see the avaiable VCs matching the selected DID and until they are valid.
        You can drag and drop new VCs in this window to add them.
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
            <div class="RightColumn">
                <div class="VClist">
                    {#key VCs.length}
                    {#each VCs as vc}
                        <div class="vc" style="padding-left: 5%;">
                            {vc.type[vc.type.length-1]}<br/>
                            Valid until: {vc.validUntil.substring(0,10)}<br/>
                            <button on:click={() => removeVC(vc)}>Remove VC</button>
                        </div> <br/>
                    {/each} 
                    {/key}
                </div>
                <button class="refreshButton" on:click={() => refreshVCs()}>Refresh VCs</button>
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

    .RightColumn {
        max-width: 80%;
        text-align: left;
        padding-left: 10%;
        padding-right: 10%;
    }

    .VClist {
        text-align: left;
        border: 1px solid black;
        border-color: black;
        max-width: 90%;
    }

    .DIDlist {
        max-width: 90%;
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
        max-width: 80%;
        padding: 2%;
        max-height: 5%;
        border-radius: 0.7em;
    }
    .button:hover {opacity: 1}

    .refreshButton {
        background-color: #1d005f;
        border: none;
        color: white;
        text-align: center;
        font-size: 16px;
        margin: 4px 2px;
        opacity: 0.6;
        transition: 0.3s;
        display: inline-block;
        text-decoration: none;
        cursor: pointer;
        width: 90%;
        height: 2em;
        max-height: 2em;
        border-radius: 0.7em;
    }
    .refreshButton:hover {opacity: 1}

    .vc {
        padding-top: 0.2em;
    }

    .tutorial {
        text-align: left;
        padding-top: 5%;
        padding-left: 5%;
        padding-right: 5%;
    }
</style>