package com.vulnerabilities.app.http;

import android.os.Bundle;
import android.util.Log;
import androidx.appcompat.app.AppCompatActivity;
import com.vulnerabilities.app.databinding.ActivityHttpBinding;
import com.vulnerabilities.app.util.HttpUtil;
import com.vulnerabilities.app.util.OtherUtil;

public class HttpActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // 1. 绑定布局
        ActivityHttpBinding httpBinding = ActivityHttpBinding.inflate(getLayoutInflater());
        setContentView(httpBinding.getRoot());

        // 2. 点击事件绑定
        httpBinding.btnHttpGet.setOnClickListener(v -> {
            Log.d("HttpActivity", "进入btnHttpGet");
            new Thread(() -> {
                try {
                    String user = "a7cc";
                    String pwd = "123456";
                    String result = HttpUtil.ConnGet("https://www.example.com/api?user="+user+"&pwd="+pwd, user, pwd);
                    runOnUiThread(() -> {
                        httpBinding.webResultHttp.loadDataWithBaseURL(
                                null,
                                result,
                                "text/html",
                                "UTF-8",
                                null
                        );
                    });
                } catch (Exception e) {
                    Log.e("HttpActivity", "ConnGet请求失败", e);
                    runOnUiThread(() -> {
                        httpBinding.webResultHttp.loadDataWithBaseURL(null,
                                "ConnGet请求异常",
                                "text/html",
                                "UTF-8",
                                null
                                );
                    });
                }
            }).start();
        });

        httpBinding.btnHttpPost.setOnClickListener(v -> {
            Log.d("HttpActivity", "进入btnHttpPost");
            new Thread(() -> {
                try {
                    String user = "a7cc";
                    String pwd = "123456";
                    String aid = OtherUtil.getAndroidId(this);
                    String result = HttpUtil.ConnPost("https://www.example.com/api?user="+user+"&pwd="+pwd, user, pwd, aid);
                    runOnUiThread(() -> {
                        httpBinding.webResultHttp.loadDataWithBaseURL(
                                null,
                                result,
                                "text/html",
                                "UTF-8",
                                null
                        );
                    });
                } catch (Exception e) {
                    Log.e("HttpActivity", "ConnPost请求失败", e);
                    runOnUiThread(() -> {
                        httpBinding.webResultHttp.loadDataWithBaseURL(null,
                                "ConnPost请求异常",
                                "text/html",
                                "UTF-8",
                                null
                        );
                    });
                }
            }).start();
        });
        httpBinding.btnOkhttp3Get.setOnClickListener(v -> {
            Log.d("HttpActivity", "进入btnOkhttp3Get");
            new Thread(() -> {
                try {
                    String user = "a7cc";
                    String pwd = "123456";
                    String result = HttpUtil.OkGet("https://www.example.com/api?user="+user+"&pwd="+pwd, user, pwd);
                    runOnUiThread(() -> {
                        httpBinding.webResultHttp.loadDataWithBaseURL(
                                null,
                                result,
                                "text/html",
                                "UTF-8",
                                null
                        );
                    });
                } catch (Exception e) {
                    Log.e("HttpActivity", "http3Get请求失败", e);
                    runOnUiThread(() -> {
                        httpBinding.webResultHttp.loadDataWithBaseURL(null,
                                "http3Get请求异常",
                                "text/html",
                                "UTF-8",
                                null
                        );
                    });
                }
            }).start();
        });
        httpBinding.btnOkhttp3Post.setOnClickListener(v -> {
            Log.d("HttpActivity", "进入btnOkhttp3Post");
            new Thread(() -> {
                try {
                    String user = "a7cc";
                    String pwd = "123456";
                    String aid = OtherUtil.getAndroidId(this);
                    String result = HttpUtil.OkPost("https://www.example.com/api?user="+user+"&pwd="+pwd, user, pwd, aid);
                    runOnUiThread(() -> {
                        httpBinding.webResultHttp.loadDataWithBaseURL(
                                null,
                                result,
                                "text/html",
                                "UTF-8",
                                null
                        );
                    });
                } catch (Exception e) {
                    Log.e("HttpActivity", "http3Post请求失败", e);
                    runOnUiThread(() -> {
                        httpBinding.webResultHttp.loadDataWithBaseURL(null,
                                "http3Post请求异常",
                                "text/html",
                                "UTF-8",
                                null
                        );
                    });
                }
            }).start();
        });
    }
}

