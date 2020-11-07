<?php
function EncryptData($source) {
    $fp=fopen("certK.pem","r");
    $pub_key=fread($fp,8192);
    fclose($fp);

    openssl_public_encrypt($source,$crypttext, $pub_key );
    // echo($crypttext);

    // private key decryption
    $fp=fopen("keyK.pem","r");
    $priv_key=fread($fp,8192);
    fclose($fp);

    openssl_private_decrypt($crypttext,$decr, $priv_key );
    // echo($crypttext);
    echo($decr);
    // return(base64_encode($crypttext));
}

echo(EncryptData("so long over the rainbow. That is a song that is ringing in my head."));