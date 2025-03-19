package com.kekwy.cfn.example;

import lombok.AllArgsConstructor;
import lombok.Data;

import java.util.Collection;

@Data
@AllArgsConstructor
public class Resource {
    private int cpu;
    private int memory;
    private Collection<String> tags;

    @SuppressWarnings("BooleanMethodIsAlwaysInverted")
    public boolean check(Resource requiredResource) {
        return cpu >= requiredResource.cpu && memory >= requiredResource.memory
                && tags.containsAll(requiredResource.tags);
    }

    public void allocate(Resource requiredResource) {
        // TODO: tag 所表示的设备是否为独占，或也需要对其负载进行量化
        cpu -= requiredResource.cpu;
        memory -= requiredResource.memory;
    }


    public void recycle(Resource recycledResource) {
        cpu += recycledResource.cpu;
        memory += recycledResource.memory;
    }
}
