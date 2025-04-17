package com.github9triver.cfn;

import lombok.Data;

import java.util.Collection;

@Data
public class HeadNodeState {
    private String id;
    private Resource totalResources;
    private Collection<String> workers;
    private Collection<String> neighbours;
}
