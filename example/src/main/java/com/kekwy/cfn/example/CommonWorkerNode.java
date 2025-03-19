package com.kekwy.cfn.example;

import lombok.Getter;
import lombok.NoArgsConstructor;

import java.util.Collection;

@NoArgsConstructor
public class CommonWorkerNode implements WorkerNode {

    private String id;

    @Getter
    private Resource resource;

    public CommonWorkerNode(int cpu, int mem, Collection<String> tags) {
        resource = new Resource(cpu, mem, tags);
    }

    @Override
    public void setId(String id) {
        this.id = id;
    }

    @Override
    public String getId() {
        return id;
    }

    @Override
    public void start() {

    }
}
