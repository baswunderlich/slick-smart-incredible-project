chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {
  if (request.action === "readEmail") {
    const parentDiv = document.getElementById('id6');
    if (parentDiv && parentDiv.children.length > 0) {
        const thirdChild = parentDiv.children[2]; // Zugriff auf das dritte Kind
        if (thirdChild.shadowRoot) {
            const firstDivInShadow = thirdChild.shadowRoot.querySelector('div');
            if (firstDivInShadow) {
                console.log('Erstes Div im Shadow DOM gefunden:', firstDivInShadow);

                const iframe = firstDivInShadow.querySelector('iframe'); // Suche nach dem iframe im ersten Div
                if (iframe) {
                    console.log('iframe gefunden:', iframe);
                    try {
                        const iframeDoc = iframe.contentDocument || iframe.contentWindow.document; // Zugriff auf den Dokument-Inhalt des iframe
                        const firstDiv = iframeDoc.querySelector('div'); // Suche nach dem ersten Div im iframe
                        if (firstDiv) {
                            console.log('Erstes Div im iframe gefunden:', firstDiv);
                            const childDiv = firstDiv.querySelector('div'); // Suche nach dem ersten Kind-Div des gefundenen Div
                            if (childDiv) {
                                console.log('Kind-Div des ersten Div im iframe gefunden:', childDiv);
								var mailText = childDiv.innerText;
								var reversedText = reverseString(mailText);
								childDiv.innerText += "\n\n\n\nReversed Text:\n " + reversedText;
                            } else {
                                console.error('Kein Kind-Div im ersten Div des iframe gefunden.');
                            }
                        } else {
                            console.error('Kein erstes Div im iframe gefunden.');
                        }
                    } catch (error) {
                        console.error('Fehler beim Zugriff auf den Inhalt des iframe:', error);
                    }
                } else {
                    console.error('Kein iframe im Shadow DOM gefunden.');
                }
            } else {
                console.error('Kein erstes Div im Shadow DOM gefunden.');
            }
        } else {
            console.error('Kein Shadow DOM verfügbar oder das dritte Kind ist kein benutzerdefiniertes Element mit einem Shadow DOM.');
        }
    } else {
        if (parentDiv) {
            console.error('Nicht genug Kinder in #id6 oder drittes Kind existiert nicht.');
        } else {
            console.error('Übergeordnetes Element mit der ID "id6" wurde nicht gefunden.');
        }
    }
  }
  /* if (request.action === "readEmail") {
	  
	      const parentDiv = document.getElementById('id6');
    if (parentDiv && parentDiv.children.length > 0) {
        const thirdChild = parentDiv.children[2]; // Zugriff auf das dritte Kind
        if (thirdChild.shadowRoot) {
            const firstDivInShadow = thirdChild.shadowRoot.querySelector('div');
            if (firstDivInShadow) {
                console.log('Erstes Div im Shadow DOM gefunden:', firstDivInShadow);

                const iframe = firstDivInShadow.querySelector('iframe'); // Erstes Div des ersten Div
                if (iframe) {
                    console.log('iframe gefunden:', firstDivChild);
                } else {
                    console.error('Kein iframe im Shadow DOM gefunden.');
                }
            } else {
                console.error('Kein erstes Div im Shadow DOM gefunden.');
            }
        } else {
            console.error('Kein Shadow DOM verfügbar oder das dritte Kind ist kein benutzerdefiniertes Element mit einem Shadow DOM.');
        }
    } else {
        if (parentDiv) {
            console.error('Nicht genug Kinder in #id6 oder drittes Kind existiert nicht.');
        } else {
            console.error('Übergeordnetes Element mit der ID "id6" wurde nicht gefunden.');
        }
    }
  } */
});

function reverseString(str) {
    return str.split("").reverse().join("");
}
