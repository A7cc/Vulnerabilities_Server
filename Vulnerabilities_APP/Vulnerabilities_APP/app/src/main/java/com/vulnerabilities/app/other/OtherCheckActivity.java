package com.vulnerabilities.app.other;

import android.os.Bundle;
import android.util.Log;
import android.widget.Toast;
import androidx.appcompat.app.AppCompatActivity;
import com.vulnerabilities.app.databinding.ActivityOtherBinding;
import com.vulnerabilities.app.util.OtherUtil;

public class OtherCheckActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // 1. 绑定布局
        ActivityOtherBinding otherBinding = ActivityOtherBinding.inflate(getLayoutInflater());
        setContentView(otherBinding.getRoot());

        // 2. 点击事件绑定
        otherBinding.btnCheckPhone0.setOnClickListener(v -> {
            Log.d("OtherCheckActivity", "进入 btnCheckPhone0 检测");
            String brandName = OtherUtil.getBrand();
            String modelName = OtherUtil.getModel();
            String manufacturer = OtherUtil.getManufacturer();
            Log.d("OtherCheckActivity", "<UNK> btnCheckPhone0 <UNK>"+brandName+modelName+manufacturer);
            if (brandName.equals("a7cc") && manufacturer.equals("a7") && modelName.equals("demo")) {
                showToast("a7cc牌手机🐮!");
            } else {
                showToast("不合规手机!!!");
            }
        });
        otherBinding.btnCheckOperator0.setOnClickListener(v -> {
            Log.d("OtherCheckActivity", "进入btnCheckOperator0检测");
            String operator = OtherUtil.getOperator(this);
            Log.d("OtherCheckActivity", "<UNK>operator<UNK>" + operator);
            if (operator.equals("a7cc")) {
                showToast("还得是a7cc运营商信号好!");
            } else {
                showToast("能不能别用这个运营商!!!");
            }
        });
        otherBinding.btnCheckNetwork0.setOnClickListener(v -> {
            Log.d("OtherCheckActivity", "进入btnCheckNetwork0检测");
            String networkType = OtherUtil.getNetworkType(this);
            Log.d("OtherCheckActivity", "<UNK>networkType<UNK>" + networkType);

            if (networkType.equals("a7cc")) {
                showToast("这个网络速度超快!");
            } else {
                showToast("不合规的网络!!!");
            }
        });
        otherBinding.btnCheckTimezone0.setOnClickListener(v -> {
            Log.d("OtherCheckActivity", "进入btnCheckTimezone0检测");
            String getTimeZone = OtherUtil.getTimeZone();
            Log.d("OtherCheckActivity", "<UNK>getTimeZone<UNK>" + getTimeZone);
            if (getTimeZone.equals("Antarctica/Vostok") ) {
                showToast("哥们!能不能带我也去一下南极!");
            } else {
                showToast("这是什么地区？？？");
            }
        });
    }
    public void showToast(String ok) {
        Toast.makeText(this, ok, Toast.LENGTH_SHORT).show();
    }
}
