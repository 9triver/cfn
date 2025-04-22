package com.github9triver.cfn.model;

import lombok.Data;

@Data
public class Resource {

    public static final String[] RESOURCE_TYPES = {
            "cpu", "memory"
    };

    @Data
    public static class CPU {
        public static final int CORE_m = 1;
        private int cores;
    }

    public static class Memory {
//        public static final int
    }

}
