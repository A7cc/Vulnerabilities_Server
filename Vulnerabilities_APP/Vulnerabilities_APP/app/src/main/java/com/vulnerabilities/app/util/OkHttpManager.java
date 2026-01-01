package com.vulnerabilities.app.util;

import android.util.Log;
import okhttp3.OkHttpClient;
import okhttp3.Request;

import java.util.concurrent.TimeUnit;

public class OkHttpManager {

    public static final OkHttpClient client = new OkHttpClient.Builder()
            .connectTimeout(5, TimeUnit.SECONDS)
            .readTimeout(5, TimeUnit.SECONDS)
            .writeTimeout(5, TimeUnit.SECONDS)
            .addInterceptor(chain -> {
                Request request = chain.request();
                Log.d("OKHTTP", "========================");
                Log.d("OKHTTP", "URL     : " + request.url());
                Log.d("OKHTTP", "Method  : " + request.method());
                Log.d("OKHTTP", "Headers : " + request.headers());
                return chain.proceed(request);
            })

            .build();
}
