package com.vulnerabilities.app.rootcheck;

import android.os.Bundle;
import android.util.Log;
import android.widget.Toast;
import androidx.appcompat.app.AppCompatActivity;
import com.vulnerabilities.app.util.RootCheckUtil;
import com.vulnerabilities.app.databinding.ActivityRootBinding;

import static com.vulnerabilities.app.jninative.JniNative.nativeCheckRoot;

public class RootCheckActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // 1. 绑定布局
        ActivityRootBinding rootBinding = ActivityRootBinding.inflate(getLayoutInflater());
        setContentView(rootBinding.getRoot());

        // 2. 点击事件绑定
        // Java层su文件检测
        rootBinding.btnCheckRoot0.setOnClickListener(v -> {
            Log.d("RootCheckActivity", "进入checkSuFile检测");
            boolean ok = RootCheckUtil.checkSuFile();
            showToast(ok);
        });
        // Java层命令执行检测
        rootBinding.btnCheckRoot1.setOnClickListener(v -> {
            Log.d("RootCheckActivity", "进入checkSuExec检测");
            boolean ok = RootCheckUtil.checkSuExec();
            showToast(ok);
        });
        // SystemProperties 反射检测
        rootBinding.btnCheckRoot2.setOnClickListener(v -> {
            Log.d("RootCheckActivity", "进入checkSystemProperty检测");
            boolean ok = RootCheckUtil.checkSystemProperty();
            showToast(ok);
        });
        // Native 层检测
        rootBinding.btnCheckRoot3.setOnClickListener(v -> {
            Log.d("RootCheckActivity", "nativeCheckRoot");
            boolean ok = nativeCheckRoot();
            showToast(ok);
        });
        // 混淆检测
        rootBinding.btnCheckRoot4.setOnClickListener(v -> {
            Log.d("RootCheckActivity", "checkConfusing");
            boolean ok = RootCheckUtil.checkConfusing();
            showToast(ok);
        });
    }
    public void showToast(boolean ok) {
        Toast.makeText(this, ok? "设备已ROOT": "设备未ROOT", Toast.LENGTH_LONG).show();
    }
}

