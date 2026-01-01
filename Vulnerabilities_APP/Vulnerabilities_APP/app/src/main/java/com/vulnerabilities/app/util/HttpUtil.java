package com.vulnerabilities.app.util;

import android.util.Log;
import okhttp3.MediaType;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;

import javax.net.ssl.HttpsURLConnection;
import java.io.*;
import java.net.URL;
import java.nio.charset.StandardCharsets;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

public class HttpUtil {
    private static HttpsURLConnection getHttpsURLConnection(URL url, String method, Map<String, String> headerMap, String body) throws IOException {
        HttpsURLConnection conn = (HttpsURLConnection) url.openConnection();
        // 请求方式
        conn.setRequestMethod(method);
        // 超时
        conn.setConnectTimeout(5000);
        conn.setReadTimeout(5000);
        // 请求头
        conn.setRequestProperty("User-Agent", "Android-App");
        for (String key : headerMap.keySet()) {
            conn.setRequestProperty(key, headerMap.get(key));
        }
        if (method.equals("GET")) {
            // 发起连接
            conn.connect();
        } else {
            OutputStream os = conn.getOutputStream();
            os.write(body.getBytes(StandardCharsets.UTF_8));
            os.flush();
            os.close();
        }

        return conn;
    }
    private static String getConnResponse(HttpsURLConnection conn) throws IOException {
        // 响应码
        int code = conn.getResponseCode();
        InputStream is = (code == 200)? conn.getInputStream() : conn.getErrorStream();
        // 读取响应
        BufferedReader reader = new BufferedReader(new InputStreamReader(is));
        StringBuilder sb = new StringBuilder();
        String line;
        while ((line = reader.readLine()) != null) {
            sb.append(line);
        }
        reader.close();
        conn.disconnect();
        return sb.toString();
    }

    private static String getHttp3Response(Request request) throws IOException {
        Response response = OkHttpManager.client.newCall(request).execute();
        if (!response.isSuccessful()) {
            throw new IOException("Unexpected code " + response);
        }
        if (response.body() != null) {
            Log.d("OKHTTP", "+=======================");
            return response.body().string();
        }

        return "请求异常";
    }

    public static String ConnGet(String Url, String user, String pwd) throws Exception {
        String nonce = UUID.randomUUID().toString().replace("-", "");
        String timestamp = String.valueOf(System.currentTimeMillis() / 1000);
        pwd = CryptoUtil.c(pwd+timestamp, 1) ;
        Map<String, String> hMap = new HashMap<>();
        hMap.put("Nonce", nonce);
        String params = "user=" + user + "&pwd=" + pwd + "&timestamp=" + timestamp;
        hMap.put("Accept", "application/json");
        if (Url.contains("?")) {
            Url = Url + "&" + params;
        } else {
            Url = Url + "?" + params;
        }
        URL url = new URL(Url);
        String sign = CryptoUtil.c(nonce.substring(8)+params+ keystore, 3);
        hMap.put("Sign", sign);
        HttpsURLConnection conn = getHttpsURLConnection(url, "GET", hMap, "");
        return getConnResponse(conn);
    }

    public static String ConnPost(String Url,String user, String pwd, String aid) throws Exception {
        String nonce = UUID.randomUUID().toString().replace("-", "").substring(10);
        Map<String, String> hMap = new HashMap<>();
        hMap.put("Nonce", nonce);
        String timestamp = String.valueOf(System.currentTimeMillis() / 1000);
        hMap.put("Content-Type", "application/x-www-form-urlencoded");
        URL url = new URL(Url);

        String body = CryptoUtil.h("user="+user+"&pwd="+pwd+"&timestamp="+timestamp+"&id=", keystore);

        hMap.put("Aid", aid);
        hMap.put("Sign", CryptoUtil.c(body+aid+nonce.substring(10), 4));
        hMap.put("Timestamp", timestamp+nonce.substring(5));
        HttpsURLConnection conn = getHttpsURLConnection(url, "POST", hMap, body);
        return getConnResponse(conn);
    }

    public static String OkGet(String url, String user, String pwd) throws Exception {
        String nonce = UUID.randomUUID().toString().replace("-", "");
        String timestamp = String.valueOf(System.currentTimeMillis() / 1000);
        pwd = CryptoUtil.c(pwd + timestamp, 0);
        String params = "user=" + user + "&pwd=" + pwd + "&timestamp=" + timestamp;
        if (url.contains("?")) {
            url += "&" + params;
        } else {
            url += "?" + params;
        }
        String sign = CryptoUtil.c(nonce.substring(8) + params + keystore, 1);

        Request request = new Request.Builder().url(url).get()
                .addHeader("User-Agent", "Android-App")
                .addHeader("Nonce", nonce)
                .addHeader("Sign", sign)
                .build();
        return getHttp3Response(request);
    }
    public static String OkPost(String url, String user, String pwd, String aid) throws Exception {

        String nonce = UUID.randomUUID().toString().replace("-", "").substring(10);
        String timestamp = String.valueOf(System.currentTimeMillis() / 1000);

        String rawBody = "user=" + user +
                "&pwd=" + pwd +
                "&timestamp=" + timestamp +
                "&id=";

        String body = CryptoUtil.h(rawBody, keystore);

        String sign = CryptoUtil.c(
                body + aid + nonce.substring(10), 4
        );

        RequestBody requestBody = RequestBody.create(
                body,
                MediaType.parse("application/x-www-form-urlencoded")
        );

        Request request = new Request.Builder()
                .url(url)
                .post(requestBody)
                .addHeader("User-Agent", "Android-App")
                .addHeader("Nonce", nonce)
                .addHeader("Aid", aid)
                .addHeader("Sign", sign)
                .addHeader("Timestamp", timestamp + nonce.substring(5))
                .build();
        return getHttp3Response(request);
    }

    static String keystore = "HsSgaUds1s27sdDuU1FvxO04w81uey2";

}
