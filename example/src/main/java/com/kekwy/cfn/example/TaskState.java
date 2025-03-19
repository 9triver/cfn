package com.kekwy.cfn.example;

import lombok.Data;

@Data
public class TaskState {
    private String id;
    private String state;
    private Resource requiredResource;
}
