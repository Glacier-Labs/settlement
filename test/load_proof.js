const fs = require('fs');

function loadProof(filename) {
    const proof = fs.readFileSync(filename);
    const proofBytes = Uint8Array.from(proof);
    const fpSize = 4 * 8;

    // proof.Ar, proof.Bs, proof.Krs
    const a0 = proofBytes.slice(fpSize * 0, fpSize * 1);
    const a1 = proofBytes.slice(fpSize * 1, fpSize * 2);
    const b00 = proofBytes.slice(fpSize * 2, fpSize * 3);
    const b01 = proofBytes.slice(fpSize * 3, fpSize * 4);
    const b10 = proofBytes.slice(fpSize * 4, fpSize * 5);
    const b11 = proofBytes.slice(fpSize * 5, fpSize * 6);
    const c0 = proofBytes.slice(fpSize * 6, fpSize * 7);
    const c1 = proofBytes.slice(fpSize * 7, fpSize * 8);

    const a = [convertUint8ArrayToBigInt(a0), convertUint8ArrayToBigInt(a1)];
    const b =     [
        [convertUint8ArrayToBigInt(b00), convertUint8ArrayToBigInt(b01)],
        [convertUint8ArrayToBigInt(b10), convertUint8ArrayToBigInt(b11)],
    ];
    const c = [convertUint8ArrayToBigInt(c0), convertUint8ArrayToBigInt(c1)];

    return [a, b, c];
}

function convertUint8ArrayToBigInt(uint8Array) {
    // Convert the Uint8Array to a hexadecimal string
    const hexString = Array.from(uint8Array, byte => byte.toString(16).padStart(2, "0")).join("");
  
    // Convert the hexadecimal string to a BigInt
    const bigint = BigInt("0x" + hexString);
  
    return bigint;
  }

  
module.exports = {
    loadProof
}