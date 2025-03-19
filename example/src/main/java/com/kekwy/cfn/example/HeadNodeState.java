package com.kekwy.cfn.example;

import lombok.Data;

import java.util.Collection;

@Data
public class HeadNodeState {
    private String id;
    private Resource totalResources;
    private Collection<String> workers;
}
