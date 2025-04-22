package com.github9triver.cfn.model;

import lombok.Data;

@Data
public class Resource {

    public static final String[] RESOURCE_TYPES = {
            "cpu", "memory"
    };

}
