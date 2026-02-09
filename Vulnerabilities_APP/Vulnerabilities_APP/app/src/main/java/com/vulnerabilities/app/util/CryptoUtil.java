package com.vulnerabilities.app.util;

import android.util.Base64;
import android.util.Log;

import javax.crypto.Cipher;
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.security.*;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;

public class CryptoUtil {
    static String[] ct = new String[]{ "MD5", "SHA-1", "SHA-256", "SHA-384", "SHA-512"};
    static String[] hm = new String[]{ "HmacMD5", "HmacSHA1", "HmacSHA256", "HmacSHA512"};

    // base64
    public static String a(String input) {
        Log.d("CryptoUtil", "call a!");
        // 将字符串转换成字节数组，再进行Base64编码
        byte[] bytes = input.getBytes(StandardCharsets.UTF_8);
        return Base64.encodeToString(bytes, Base64.NO_WRAP); // NO_WRAP 去掉换行
    }
    public static String b(String base64Str) {
        Log.d("CryptoUtil", "call b!");
        // Base64编码
        byte[] bytes = Base64.decode(base64Str, Base64.NO_WRAP);
        return new String(bytes);
    }
    // md5/sha256
    public static String c(String input, int num) {
        Log.d("CryptoUtil", "call c!");
        try {
            MessageDigest md = MessageDigest.getInstance(ct[num]);
            byte[] bytes = md.digest(input.getBytes(StandardCharsets.UTF_8));
            return bytesToHex(bytes);
        } catch (Exception e) {
            Log.e("CryptoUtil", "c :" + e);
            return "";
        }
    }
    // DES 加密
    public static String d(String data, String key) {
        Log.d("CryptoUtil", "call d!");
        try {
            Key desKey = new SecretKeySpec(key.getBytes(StandardCharsets.UTF_8), "DES");
            Cipher cipher = Cipher.getInstance("DES/ECB/PKCS5Padding");
            cipher.init(Cipher.ENCRYPT_MODE, desKey);
            byte[] encrypted = cipher.doFinal(data.getBytes(StandardCharsets.UTF_8));
            return bytesToHex(encrypted);
        } catch (Exception e) {
            Log.e("CryptoUtil", "d :" + e);
            return "null";
        }
    }
    // DES 解密
    public static String e(String hexData, String key) {
        Log.d("CryptoUtil", "call e!");
        try {
            byte[] bytes = hexStringToByteArray(hexData);
            Key desKey = new SecretKeySpec(key.getBytes(StandardCharsets.UTF_8), "DES");
            Cipher cipher = Cipher.getInstance("DES/ECB/PKCS5Padding");
            cipher.init(Cipher.DECRYPT_MODE, desKey);
            byte[] decrypted = cipher.doFinal(bytes);
            return new String(decrypted);
        } catch (Exception e) {
            Log.e("CryptoUtil", "e :" + e);
            return "null";
        }
    }
    // 3des
    public static String f(String data, String key) {
        Log.d("CryptoUtil", "call f!");
        try {
            SecretKeySpec keySpec = new SecretKeySpec(key.getBytes(StandardCharsets.UTF_8), "DESede");
            Cipher cipher = Cipher.getInstance("DESede/ECB/PKCS5Padding");
            cipher.init(Cipher.ENCRYPT_MODE, keySpec);
            byte[] encrypted = cipher.doFinal(data.getBytes(StandardCharsets.UTF_8));
            return bytesToHex(encrypted);
        } catch (Exception e) {
            Log.e("CryptoUtil", "f :" + e);
            return "null";
        }
    }
    public static String g(String hexData, String key) {
        Log.d("CryptoUtil", "call g!");
        try {
            byte[] bytes = hexStringToByteArray(hexData);
            SecretKeySpec keySpec = new SecretKeySpec(key.getBytes(StandardCharsets.UTF_8), "DESede");
            Cipher cipher = Cipher.getInstance("DESede/ECB/PKCS5Padding");
            cipher.init(Cipher.DECRYPT_MODE, keySpec);
            byte[] decrypted = cipher.doFinal(bytes);
            return new String(decrypted);
        } catch (Exception e) {
            Log.e("CryptoUtil", "g :" + e);
            return "null";
        }
    }
    // aes
    private static SecretKeySpec deriveKey(String key) {
        byte[] keyBytes = new byte[16]; // 默认取16位
        byte[] userBytes = key.getBytes(StandardCharsets.UTF_8);
        // 将用户输入的 key 复制到 keyBytes 中，不足补0，多余截断
        System.arraycopy(userBytes, 0, keyBytes, 0, Math.min(userBytes.length, keyBytes.length));
        return new SecretKeySpec(keyBytes, "AES");
    }
    public static String h(String data, String key) {
        Log.d("CryptoUtil", "call h!");
        try {
            SecretKeySpec keySpec = deriveKey(key);
            Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
            cipher.init(Cipher.ENCRYPT_MODE, keySpec);
            byte[] encrypted = cipher.doFinal(data.getBytes(StandardCharsets.UTF_8));
            return bytesToHex(encrypted);
        } catch (Exception e) {
            Log.e("CryptoUtil", "h :" + e);
            return "null";
        }
    }

    public static String i(String hexData, String key) {
        Log.d("CryptoUtil", "call i!");
        try {
            byte[] bytes = hexStringToByteArray(hexData);
            SecretKeySpec keySpec = deriveKey(key);
            Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
            cipher.init(Cipher.DECRYPT_MODE, keySpec);
            byte[] decrypted = cipher.doFinal(bytes);
            return new String(decrypted);
        } catch (Exception e) {
            Log.e("CryptoUtil", "i :" + e);
            return "null";
        }
    }

    // 公钥加密
    public static String j(String plainText, String base64PublicKey) throws Exception {
        Log.d("CryptoUtil", "call j!");
        try {
            if (plainText == null || base64PublicKey == null) return "";
            Cipher cipher = Cipher.getInstance("RSA/ECB/PKCS1Padding");
            byte[] keyBytes = Base64.decode(base64PublicKey, Base64.NO_WRAP);
            X509EncodedKeySpec spec = new X509EncodedKeySpec(keyBytes);
            KeyFactory kf = KeyFactory.getInstance("RSA");
            cipher.init(Cipher.ENCRYPT_MODE, kf.generatePublic(spec));
            byte[] encryptedBytes = cipher.doFinal(plainText.getBytes(StandardCharsets.UTF_8));
            return Base64.encodeToString(encryptedBytes, Base64.NO_WRAP);
        } catch (Exception e) {
            Log.e("CryptoUtil", "j :" + e);
            return "null";
        }
    }

    // 私钥解密
    public static String k(String cipherText, String base64PrivateKey) throws Exception {
        Log.d("CryptoUtil", "call k!");
        try {
            if (cipherText == null || base64PrivateKey == null) return "";
            Cipher cipher = Cipher.getInstance("RSA/ECB/PKCS1Padding");
            byte[] keyBytes = Base64.decode(base64PrivateKey, Base64.NO_WRAP);
            PKCS8EncodedKeySpec spec = new PKCS8EncodedKeySpec(keyBytes);
            KeyFactory kf = KeyFactory.getInstance("RSA");
            cipher.init(Cipher.DECRYPT_MODE, kf.generatePrivate(spec));
            byte[] bytes = Base64.decode(cipherText, Base64.NO_WRAP);
            byte[] decrypted = cipher.doFinal(bytes);
            return new String(decrypted, StandardCharsets.UTF_8);
        }  catch (Exception e) {
            Log.e("CryptoUtil", "k :" + e);
            return "null";
        }
    }

    // HMAC
    public static String l(String cipherText, String key, int num) throws Exception {
        Log.d("CryptoUtil", "call l!");
        try {
            if (num < 0 || num >= hm.length) return "错误：算法索引越界";
            Mac mac = Mac.getInstance(hm[num]);
            byte[] keyBytes = (key != null) ? key.getBytes(StandardCharsets.UTF_8) : new byte[0];
            SecretKeySpec secretKey = new SecretKeySpec(keyBytes, hm[num]);
            mac.init(secretKey);
            byte[] rawHmac = mac.doFinal(cipherText.getBytes(StandardCharsets.UTF_8));
            return Base64.encodeToString(rawHmac, Base64.NO_WRAP);
        } catch (Exception e) {
            Log.e("CryptoUtil", "l :" + e);
            return "null";
        }
    }

    // RSA+SHA256签名
    public static String s(String cipherText, String base64PrivateKey) throws Exception {
        Log.d("CryptoUtil", "call s!");
        try {
            byte[] keyBytes = Base64.decode(base64PrivateKey, Base64.NO_WRAP);
            PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(keyBytes);
            KeyFactory keyFactory = KeyFactory.getInstance("RSA");
            PrivateKey privateKey = keyFactory.generatePrivate(keySpec);
            // 签名
            Signature signature = Signature.getInstance("SHA256withRSA");
            signature.initSign(privateKey);
            signature.update(cipherText.getBytes(StandardCharsets.UTF_8));
            byte[] sign = signature.sign();
            // 返回base64字符串
            return Base64.encodeToString(sign, Base64.NO_WRAP);
        } catch (Exception e) {
            Log.e("CryptoUtil", "s :" + e);
            return "null";
        }
    }

    // 将字节数组转换为十六进制字符串
    private static String bytesToHex(byte[] bytes) {
        StringBuilder sb = new StringBuilder();
        for (byte b : bytes) {
            sb.append(String.format("%02X", b));
        }
        return sb.toString();
    }
    // 将十六进制字符串转换为字节数组
    private static byte[] hexStringToByteArray(String s) {
        int len = s.length();
        byte[] data = new byte[len / 2];
        for (int i = 0; i < len; i += 2) {
            data[i / 2] = (byte) ((Character.digit(s.charAt(i), 16) << 4)
                    + Character.digit(s.charAt(i+1), 16));
        }
        return data;
    }
}
