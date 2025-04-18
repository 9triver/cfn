package com.github9triver.cfn.config;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;

@Data
@ConfigurationProperties(prefix = "cfn.node")
public class NodeProperties {
    private boolean head;
}
