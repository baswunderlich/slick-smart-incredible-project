from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_OAEP
 
class KeyPair:
    pubKey: str
    privKey: str

    def __init__(self, pubKey, privKey):
        self.pubKey = pubKey
        self.privKey = privKey 

def gen_keys() -> KeyPair:
    keyPair = RSA.generate(3072)
    
    pubKey = keyPair.publickey()
    # print(f"Public key:  (n={hex(pubKey.n)}, e={hex(pubKey.e)})")
    pubKeyPEM = pubKey.exportKey()
    # print(pubKeyPEM.decode('ascii'))
    
    # print(f"Private key: (n={hex(pubKey.n)}, d={hex(keyPair.d)})")
    privKeyPEM = keyPair.exportKey()
    # print(privKeyPEM.decode('ascii'))
    return KeyPair(pubKey=pubKeyPEM, privKey=privKeyPEM)
 
def encrypt(pubKey: str, message: str) -> str:
    #encryption
    encryptor = PKCS1_OAEP.new(pubKey)
    encrypted = encryptor.encrypt(bytes(message, encoding="utf-8"))
    encrypted

def decrypt(privKey: str, message: str) -> str:
    #decryption
    decryptor = PKCS1_OAEP.new(privKey)
    decrypted = decryptor.decrypt(bytes(message, encoding="utf-8"))
    return decrypted