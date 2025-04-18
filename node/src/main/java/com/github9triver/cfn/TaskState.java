package com.github9triver.cfn;

import lombok.Data;

@Data
public class TaskState {
    private String id;
    private String state;
    private Resource requiredResource;
}
