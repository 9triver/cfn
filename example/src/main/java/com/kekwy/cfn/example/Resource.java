package com.kekwy.cfn.example;

import lombok.AllArgsConstructor;
import lombok.Data;

import java.util.Collection;

@Data
@AllArgsConstructor
public class Resource {
    private int cpu;
    private int memory;
    Collection<String> tags;
}
