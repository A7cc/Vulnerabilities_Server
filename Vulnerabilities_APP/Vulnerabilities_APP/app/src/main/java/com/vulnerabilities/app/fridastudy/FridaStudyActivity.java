package com.vulnerabilities.app.fridastudy;

import android.os.Bundle;
import android.text.TextUtils;
import android.util.Log;
import android.widget.Toast;
import androidx.appcompat.app.AppCompatActivity;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;
import com.vulnerabilities.app.databinding.ActivityFridastudyBinding;

import java.util.*;
import java.util.regex.Pattern;

// TODO: 需要阅读一下，然后flag需要全面修改
public class FridaStudyActivity extends AppCompatActivity {
    static {
        System.loadLibrary("fridastudy");
    }
    // 当前关卡状态
    private int currentLevel = 1;
    // ViewBinding
    private ActivityFridastudyBinding fridastudyBinding;
    // 设置关卡数量
    private static final int LEVEL_NUM = 10;
    // IPV4 正则表达式
    private static final Pattern IPV4_PATTERN = Pattern.compile(
            "^(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)\\." +
                    "(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)\\." +
                    "(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)\\." +
                    "(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)$"
    );
    // 在 Activity 中定义为成员变量
    protected static List<Level> fridastudyLevels;
    protected static FsLevelAdapter adapter;
    // 增加一个静态变量保存当前实例，解决静态调用非静态的问题
    private static FridaStudyActivity mInstance;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        mInstance = this; // 记录实例
        // 1. 初始化 ViewBinding
        Log.d("FridaStudy", "onCreate");
        fridastudyBinding = ActivityFridastudyBinding.inflate(getLayoutInflater());
        // 2. 设置布局
        setContentView(fridastudyBinding.getRoot());
        // 将RecyclerView设置成横向滑动的列表
        fridastudyBinding.levelFsView.setLayoutManager(
                // Context、横向排列、是否反向
                new LinearLayoutManager(this, RecyclerView.HORIZONTAL, false)
        );
        // 设置等级数组，并且只开放第一关
        fridastudyLevels = new ArrayList<>();
        for (int i = 1; i <= LEVEL_NUM; i++) {
            fridastudyLevels.add(new Level(i, i == 1, false)); // 只解锁第 1 关
        }
        // 在界面上显示选择的关卡和提示
        adapter = new FsLevelAdapter(fridastudyLevels, fridastudyLevel -> {
            currentLevel = fridastudyLevel.id;
            fridastudyBinding.tvFridaLevel.setText("第 "+currentLevel+ "关");
            fridastudyBinding.tvFridaTips.setText("这关要hook哪个函数呢？");
            // 提示选择的关卡
            Toast.makeText(this, "选择第 " + currentLevel + " 关", Toast.LENGTH_SHORT).show();
        });
        fridastudyBinding.levelFsView.setAdapter(adapter);

        // Submit按钮响应
        fridastudyBinding.btnFs.setOnClickListener(v -> {
            Log.d("FridaStudy", "onCreate");
            Map<String, String> params = new HashMap<>();
            params.put("text", fridastudyBinding.etInput.getText().toString());
            Log.d("FridaStudy", "user data: " + params.get("text"));
            String flag = "你这是什么关卡？";
            // 每个关卡通关逻辑
            switch (currentLevel) {
                case 1:
                    // 第一关：frida_hook的简单使用
                    Log.d("FridaStudy", "Level 1");
                    final int i = new Random().nextInt(100);
                    flag = LevelOne(Integer.parseInt(Objects.requireNonNull(params.get("text"))), i);
                    break;
                case 2:
                    Log.d("FridaStudy", "Level 2");
                    // 第二关：主动调用静态函数
                    // LevelTwo(params.get("text"));
                    flag = "可以尝试hook LevelTwo函数";
                    break;
                case 3:
                    // 第三关：主动设置类静态值
                    Log.d("FridaStudy", "Level 3");
                    if (FsCommon.code == 1024) {
                        flag = GenerateFlag("03_"+params.get("text"));
                    }else {
                        flag = "可以尝试hook FsCommon.code变量";
                    }
                    break;
                case 4:
                    Log.d("FridaStudy", "Level 4");
                    // 第四关：主动调用动态函数
                    flag = "可以尝试hook LevelFour函数";
                    break;
                case 5:
                    Log.d("FridaStudy", "Level 5");
                    // 第五关：主动调用动态函数2
                    flag = "可以尝试hook LevelFour函数";
                    break;
                case 6:
                    Log.d("FridaStudy", "Level 6");
                    FsCommon fscTmp = LevelSix(new FsCommon(1024, params.get("text")));
                    if (fscTmp != null && fscTmp.num == 1024 && Objects.equals(fscTmp.str, "a7cc")) {
                        flag = GenerateFlag("06_"+params.get("text"));
                    } else {
                        flag = "行不行啊？if都进不去";
                    }
                    break;
                case 7:
                    Log.d("FridaStudy", "Level 7");
                    // 第七关：主动调用动态函数传递特殊类型
                    flag = "可以尝试hook LevelSeven函数";
                    break;
                case 8:
                    Log.d("FridaStudy", "Level 8");
                    if (IPV4_PATTERN.matcher(Objects.requireNonNull(params.get("text"))).matches() && Objects.equals(LevelEight(params.get("text")), "a7cc")) {
                        flag = GenerateFlag("08"+params.get("text"));
                    } else {
                        flag = "行不行啊？if都进不去";
                    }
                    break;
                case 9:
                    Log.d("FridaStudy", "Level 9");
                    flag = LevelNine(params.get("text"));
                    break;
                case 10:
                    Log.d("FridaStudy", "Level 10");
                    flag = LevelTen(params.get("text"));
                    break;
                default:
                    flag = "你这是什么关卡？";
            }
            // 对结果的判断
            showResults(flag);
        });

    }
    // 对结果判断和显示
    public static void showResultsStatic(String flag) {
        if (mInstance != null) {
            // 关键：切换到 UI 线程，防止 Frida 调用导致崩溃
            mInstance.runOnUiThread(() -> mInstance.showResults(flag));
        }
    }

    public void showResults(String flag) {
        Log.d("FridaStudy", "Get flag: " + flag);
        if (TextUtils.isEmpty(flag)) {
            fridastudyBinding.tvFridaTips.setText("flag获取失败");
            return;
        }
        if (currentLevel < fridastudyLevels.size() && flag.startsWith("flag{") && flag.endsWith("}")) {
            Toast.makeText(this, flag, Toast.LENGTH_SHORT).show();
            int nextIndex = currentLevel;
            // 解锁下一关 (currentLevel 对应的索引正好是下一关)
            fridastudyLevels.get(currentLevel).unlocked = true;
            // 只刷新被解锁的这一关
            adapter.notifyItemChanged(nextIndex);
        }
        fridastudyBinding.tvFridaTips.setText(flag);
    }

    // 第一关：frida_hook的简单使用
    String LevelOne(int text, int random) {
        Log.d("FridaStudy", "LevelOne");
        if (random * 20 + 10 == text) {
            return GenerateFlag("01_"+text+"_"+random);
        } else {
            return "你这个值不对啊！";
        }
    }
    // 第二关：主动调用静态函数
    public static void LevelTwo(String text) {
        Log.d("FridaStudy", "LevelTwo");
        if (Objects.equals(text, "a7cc")) {
            String flag = GenerateFlag("02_"+text);
            // 直接调用 Activity 的显示逻辑
            showResultsStatic(flag);
        }
    }
    // 第五关：主动调用动态函数2
    public void LevelFive(int text) {
        if (text == 1024) {
            try {
                String flag = GenerateFlag("05_"+text);
                showResultsStatic(flag);
            } catch (Exception e) {
                Log.e("FridaStudy", "LevelFive error", e);
            }
        }
    }
    // 第六关：
    public FsCommon LevelSix(FsCommon date) {
        date.num = 1;
        return date;
    }
    // 第七关：主动调用动态函数传递特殊类型
    public void LevelSeven(FsCommon date) {
        if (date.num == 1024 && Objects.equals(date.str, "a7cc")) {
            String flag = GenerateFlag("07_"+date.num + date.str);;
            // 直接调用 Activity 的显示逻辑
            showResultsStatic(flag);
        }
    }
    // 第八关：hook native函数修改返回值
    public static native String LevelEight(String ip);
    // 第九关：native层获取参数值
    public static native String LevelNine(String str);
    // 第十关：动态注册函数
    public static native String LevelTen(String str);
    // 获取flag
    public static native String GenerateFlag(String data);
}
