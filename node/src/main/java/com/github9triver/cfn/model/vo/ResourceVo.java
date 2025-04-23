package com.github9triver.cfn.model.vo;

import lombok.Data;

import java.io.Serial;
import java.io.Serializable;

@Data
public class ResourceVo implements Serializable {

    @Serial
    private static final long serialVersionUID = 1L;

    private final CPU cpu = new CPU();
    private final Memory memory = new Memory();

    @Data
    public static class CPU implements Serializable {
        @Serial
        private static final long serialVersionUID = 1L;

        private String cores;
    }

    @Data
    public static class Memory implements Serializable {
        @Serial
        private static final long serialVersionUID = 1L;

        private String capacity;
    }

}
