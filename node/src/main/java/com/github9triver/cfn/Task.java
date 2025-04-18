package com.github9triver.cfn;

import lombok.Getter;
import lombok.Setter;

import java.util.*;
import java.util.function.Function;

@Getter
@Setter
public class Task<T> {

    private String tid = UUID.randomUUID().toString(); // 暂时使用 uuid
    private final Function<Map<String, Object>, T> function;
    private final Map<String, Object> params;
    private final ResultWrapper resultWrapper = new ResultWrapper();
    private final Resource requiredResource;

    private final List<String> path = new ArrayList<>();

    public Task(Function<Map<String, Object>, T> function, Map<String, Object> params,
                int cpu, int mem, Collection<String> tags) {
        this.function = function;
        this.params = params;
        this.requiredResource = new Resource(cpu, mem, tags);
    }

    private String state;

    public void run() {
        T res = function.apply(params);
        setResult(res);
    }

    private class ResultWrapper {
        private volatile T value = null;
    }

    public T getResult() throws InterruptedException {
        if (resultWrapper.value == null) {
            synchronized (resultWrapper) {
                if (resultWrapper.value == null) {
                    resultWrapper.wait();
                }
            }
        }
        return resultWrapper.value;
    }

    public void setResult(T value) {
        synchronized (resultWrapper) {
            resultWrapper.value = value;
            resultWrapper.notify();
        }
    }

}
