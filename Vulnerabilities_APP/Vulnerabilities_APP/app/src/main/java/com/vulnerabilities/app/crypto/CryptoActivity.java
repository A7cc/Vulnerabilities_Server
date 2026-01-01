package com.vulnerabilities.app.crypto;

import android.os.Bundle;
import android.util.Log;
import androidx.appcompat.app.AppCompatActivity;
import com.vulnerabilities.app.databinding.ActivityCryptoBinding;
import com.vulnerabilities.app.util.CryptoUtil;


public class CryptoActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // 1. 绑定布局
        ActivityCryptoBinding cryptoBinding = ActivityCryptoBinding.inflate(getLayoutInflater());
        setContentView(cryptoBinding.getRoot());

        // 2. 点击事件绑定
        // BASE64
        cryptoBinding.btnBase64Encrypt0.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btnBase64Encrypt0");
            String result = CryptoUtil.a(cryptoBinding.etInput.getText().toString());
            cryptoBinding.tvCryptoResult.setText("Base64加密结果："+result);
        });
        cryptoBinding.btnBase64Decrypt0.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btnBase64Decrypt0");
            String result = CryptoUtil.b(cryptoBinding.etInput.getText().toString());
            cryptoBinding.tvCryptoResult.setText("Base64解密结果："+result);
        });
        // MD5
        cryptoBinding.btnMd50.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btnMd50");
            String result = CryptoUtil.c(cryptoBinding.etInput.getText().toString(), 0);
            cryptoBinding.tvCryptoResult.setText("Md5加密结果："+result);
        });
        // SHA256
        cryptoBinding.btnSha2560.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btnSha2560");
            String result = CryptoUtil.c(cryptoBinding.etInput.getText().toString(), 2);
            cryptoBinding.tvCryptoResult.setText("Sha256加密结果："+result);
        });
        // DES
        cryptoBinding.btnDesEncrypt0.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btnDesEncrypt0");
            String key = "a7cc+app";
            String result = CryptoUtil.d(cryptoBinding.etInput.getText().toString(), key);
            cryptoBinding.tvCryptoResult.setText("Des加密结果："+result);
        });
        cryptoBinding.btnDesDecrypt0.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btnDesDecrypt0");
            String key = "a7cc+app";
            String result = CryptoUtil.e(cryptoBinding.etInput.getText().toString(), key);
            cryptoBinding.tvCryptoResult.setText("Des解密结果："+result);
        });
        // 3DES
        cryptoBinding.btn3desEncrypt0.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btn3desEncrypt0");
            String key = "https://github.com/a7cc/";
            String result = CryptoUtil.f(cryptoBinding.etInput.getText().toString(), key);
            cryptoBinding.tvCryptoResult.setText("3Des加密结果："+result);
        });
        cryptoBinding.btn3desDecrypt0.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btn3desDecrypt0");
            String key = "https://github.com/a7cc/";
            String result = CryptoUtil.g(cryptoBinding.etInput.getText().toString(), key);
            cryptoBinding.tvCryptoResult.setText("3Des解密结果："+result);
        });
        // AES
        cryptoBinding.btnAesEncrypt0.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btnAesEncrypt0");
            String key = "github.com/a7cc/";
            String result = CryptoUtil.h(cryptoBinding.etInput.getText().toString(), key);
            cryptoBinding.tvCryptoResult.setText("Aes加密结果："+result);
        });
        cryptoBinding.btnAesDecrypt0.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btnAesDecrypt0");
            String key = "github.com/a7cc/";
            String result = CryptoUtil.i(cryptoBinding.etInput.getText().toString(), key);
            cryptoBinding.tvCryptoResult.setText("Aes解密结果："+result);
        });
        // RSA
        cryptoBinding.btnRsaEncrypt0.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btnRsaEncrypt0");
            try {
                String publicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp2Iz+nU49FTO0Z5Ch9JOderMPEwgZWuZOGJ3ZPqbFaoE27G88kJZ/qjAySApQ/cdeDr95K2G6CqFHuPDqAnF3MC9b/vVYsm1UOQYpre7mzCwdR+OCkd9br32575u3+kd5+DWtL3dd0QI8BiraIWBxooFgDBlRsXhzpjzL9cLHLdQhZjMS2l269ZrVmq3XPF6hLmSiE67eBX5TQTX5sfAH0NkMmMy6SOYfKsU7mus3PDNy/UBKTiQik11OtzRexme/miTDoPFAlCltFl+KY42soTSKvzHsLsRdQYoOlVHyBuAG/7ZqG25i/tKCxQJbwuESocy7WvWsRgh/4SJzJzruQIDAQAB";
                String result = CryptoUtil.j(cryptoBinding.etInput.getText().toString(), publicKey);
                cryptoBinding.tvCryptoResult.setText("Rsa加密结果：" + result);
            } catch (Exception e) {
                cryptoBinding.tvCryptoResult.setText("Rsa加密结果：err");
                throw new RuntimeException(e);
            }
        });
        cryptoBinding.btnRsaDecrypt0.setOnClickListener(v -> {
            Log.d("CryptoActivity", "进入btnRsaDecrypt0");
            try {
                String privateKey = "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCnYjP6dTj0VM7RnkKH0k516sw8TCBla5k4Yndk+psVqgTbsbzyQln+qMDJIClD9x14Ov3krYboKoUe48OoCcXcwL1v+9ViybVQ5Bimt7ubMLB1H44KR31uvfbnvm7f6R3n4Na0vd13RAjwGKtohYHGigWAMGVGxeHOmPMv1wsct1CFmMxLaXbr1mtWardc8XqEuZKITrt4FflNBNfmx8AfQ2QyYzLpI5h8qxTua6zc8M3L9QEpOJCKTXU63NF7GZ7+aJMOg8UCUKW0WX4pjjayhNIq/MewuxF1Big6VUfIG4Ab/tmobbmL+0oLFAlvC4RKhzLta9axGCH/hInMnOu5AgMBAAECggEAApNjymBgan4dmqMtUeq+E1l0IB9pxjXEK1Z0kOS8WzZ9EaZoCnBT8zPEl7O7ftTAKtf5jTnQYCU7XDDUE3NnDJEYuS7uZzFEbYS9Dl+4bthudrbwGB9mgjxBHurSyPZM5AMVOr3LOSkeXBd9F8QqLZR1JjZYcKDd6zB2WY+dF6mCHunYngcCYwFGcxU/We1zTdry4Lab0v5k5uytRZdSWEpCevo6EYn3TuM3ztIjhZn4GQTNCmui18bvvOJKreswVqjaJRmRX/0s+7fRUxaK2KQ0CjbsgQUyvbIjQFvw/lK+RojR9cszPHstbBTOfv3RjLnN6s6QmUKPh1QPMwRvMwKBgQDXi74Mm9pXQYZWPyuM+G6r6XivIVtsb73v3aaMbhl7xubsuBntN3dO3A45kN+bkIGE6GIpsw5qFicYrUx3q4Rufe3V28u1g4sEFHcJYr64TWuHiH+J3mx8cFG80r8WUHJK8yA4de9ldqsXoCqkbvfqEWUgpAiGzLcuIwQX7vxP+wKBgQDGzGzTxXv6+z9XIDVITTUzDOkgaPwOh9ggdCvbdXvlQLoxW2RLdwbSO1KvcNpKeqLkHdsr71RZHgbBHsESjf977njCCIbyy+1dXjCg+Bb3jtfhkJpAP0fqdpV21A4r0NOTAsTdtbB0ZvILxJY91TlqG9nxo8KnYp/1ihuJEqWA2wKBgESknS6Yx6z4EhcPYQgw6dXXsXZccigTYfKsrOiV+4meq1YBv0f7XQSBMgqFJ6D8ITM5amGLQ3DngyXpsqt9tNwXQJHVwIca0D/JYKtdhg1donv0LWGzsuriPjWbC/3IOs3BpaY1cLroUs5gVJQvPsaNKZ5+frSnJ0MLJbeVhUbXAoGALDrNN6yncea2Z/cFg8mRvYLw10IkWkNFBatzGoegAhiNlG+l99hKjzmdouTz3EA+v4wY+ERuOsmgbRUflbY8EmGzQlBNXfWbIPvUy+uGiJuiUlAu0X6CzJqHlIGHiZ4ThJDIJh/HW46P6ahaWPV4qa/6pHVv2Hfr6OBKUgvxAKMCgYAlh48xZBEmjAoRIJsr60h0m2LV/LdeohjddESxN8EoV1VZ53vLKQ6+MsRzkWnLFc/d2mpMa0kU4Aoiwjxpyp6/g8gZZ5Ve1eEBNg/aQ33CSyoXGNQbs1grIH6ewwqUB9PQ8u5u3uDFzC2dKmYHg5vWea+FxADE1QexlimIXZjEDg==";
                String result = CryptoUtil.k(cryptoBinding.etInput.getText().toString(), privateKey);
                cryptoBinding.tvCryptoResult.setText("Rsa解密结果：" + result);
            } catch (Exception e) {
                cryptoBinding.tvCryptoResult.setText("Rsa解密结果：err");
                throw new RuntimeException(e);
            }
        });
    }
}
