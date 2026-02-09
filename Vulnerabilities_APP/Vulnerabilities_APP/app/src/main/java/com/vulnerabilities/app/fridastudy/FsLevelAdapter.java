package com.vulnerabilities.app.fridastudy;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;
import androidx.recyclerview.widget.RecyclerView;
import com.vulnerabilities.app.R;

import java.util.List;

public class FsLevelAdapter extends RecyclerView.Adapter<FsLevelAdapter.VH> {

    private final List<Level> fridastudyLevels;
    private final OnLevelClick listener;

    public interface OnLevelClick {
        void onClick(Level fridastudyLevel);
    }

    public FsLevelAdapter(List<Level> fridastudyLevels, OnLevelClick listener) {
        this.fridastudyLevels = fridastudyLevels;
        this.listener = listener;
    }

    static class VH extends RecyclerView.ViewHolder {
        TextView tv;

        VH(View v) {
            super(v);
            tv = v.findViewById(R.id.tvLevel);
        }
    }

    // 一共有几关
    @Override
    public VH onCreateViewHolder(ViewGroup parent, int viewType) {
        View v = LayoutInflater.from(parent.getContext())
                .inflate(R.layout.fridastudy_item_level, parent, false);
        return new VH(v);
    }

    // 每一关的样子
    @Override
    public void onBindViewHolder(VH h, int pos) {
        Level l = fridastudyLevels.get(pos);
        h.tv.setText(String.valueOf(l.id));

        if (!l.unlocked) {
            h.tv.setAlpha(0.3f);
            h.tv.setEnabled(false);
        } else {
            h.tv.setAlpha(1f);
            h.tv.setEnabled(true);
        }

        h.tv.setOnClickListener(v -> listener.onClick(l));
    }

    // 用户点击的关卡
    @Override
    public int getItemCount() {
        return fridastudyLevels.size();
    }
}

