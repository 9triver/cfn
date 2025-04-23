package com.github9triver.cfn.model.dto;

import lombok.Data;

import java.math.BigDecimal;

@Data
public class ResourceDto {

    private CPU cpu = new CPU();
    private Memory memory = new Memory();

    public static final String[] RESOURCE_TYPES = {
            "cpu", "memory"
    };

    public void set(String key, BigDecimal value) {
        switch (key) {
            case "cpu":
                cpu.setCores(value);
                break;
            case "memory":
                memory.setCapacity(value);
                break;
        }
    }

    @Data
    public static class CPU {
        public static final int CORE_m = 1;
        private BigDecimal cores;
    }

    @Data
    public static class Memory {
        //        public static final int
        private BigDecimal capacity;
    }

}
