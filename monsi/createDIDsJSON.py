from cryptography.hazmat.primitives import serialization as crypto_serialization
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.backends import default_backend as crypto_default_backend
import json
import base64

class DID:
    def __init__(self, did, pubKey, privKey):
        self.did = did
        self.pubKey = pubKey
        self.privKey = privKey

class MyEncoder(json.JSONEncoder):
    def default(self, o):
        return o.__dict__    

def genKeys() -> tuple[bytes, bytes]:
    #Returns a private and a public key
    key = rsa.generate_private_key(
        backend=crypto_default_backend(),
        public_exponent=65537,
        key_size=2059
    )

    private_key = key.private_bytes(
        crypto_serialization.Encoding.PEM,
        crypto_serialization.PrivateFormat.PKCS8,
        crypto_serialization.NoEncryption()
    )

    public_key = key.public_key().public_bytes(
        crypto_serialization.Encoding.PEM,
        crypto_serialization.PublicFormat.PKCS1
    )

    return private_key, public_key



#Main

file_path = "dids.json"
 
with open(file_path, 'w') as file:
    dids = ["did:example:1", "did:example:2", "did:example:university"]
    didObjects = []
    for i in range(3):
        privKey, pubKey = genKeys()
        newDid = DID(did=dids[i], privKey=privKey.decode("utf-8"), pubKey=pubKey.decode("utf-8"))
        didObjects.append(newDid)   
    didsAsString = json.dumps(didObjects, cls=MyEncoder)
    file.write(didsAsString)

print(f"File '{file_path}' created successfully.")
