document.getElementById('readEmail').addEventListener('click', function handleButtonClick() {
    if (isProcessing) return;
    isProcessing = true;

    chrome.tabs.query({active: true, currentWindow: true}, function(tabs) {
        chrome.tabs.sendMessage(tabs[0].id, {action: "readEmail"}, function(response) {
            isProcessing = false;
        });
    });
    // Entferne den Event Listener, um doppelte Events zu vermeiden
    document.getElementById('readEmail').removeEventListener('click', handleButtonClick);
});

let isProcessing = false;

document.getElementById('readEmail').addEventListener('click', function() {

});
