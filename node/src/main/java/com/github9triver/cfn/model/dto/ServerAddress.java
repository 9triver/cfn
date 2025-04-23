package com.github9triver.cfn.model.dto;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class ServerAddress {
    private String host;
    private int port;
}
