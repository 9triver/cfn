package com.kekwy.cfn.example;

import lombok.Data;

import java.util.Collection;

@Data
public class WorkerNodeState {

    private String id;
    private Resource resource;
    private Collection<TaskState> tasks;

}
