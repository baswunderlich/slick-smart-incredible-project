import rsa
 
class KeyPair:
    pubKey: str
    privKey: str

    def __init__(self, pubKey, privKey):
        self.pubKey = pubKey
        self.privKey = privKey 

def gen_keys() -> KeyPair:
    (bob_pub, bob_priv) = rsa.newkeys(512)
    return KeyPair(pubKey=bob_pub.save_pkcs1(), privKey=bob_priv.save_pkcs1())
 
def encrypt(pubKey: str, message: str) -> str:
    #encryption
    return str(rsa.encrypt(message=bytes(message, encoding="utf-8"), pub_key=rsa.PublicKey.load_pkcs1(pubKey)))

def decrypt(privKey: str, message: str) -> str:
    #decryption   
    return str(rsa.decrypt(message=bytes(message, encoding="utf-8"), priv_key=rsa.PublicKey.load_pkcs1(privKey)))
