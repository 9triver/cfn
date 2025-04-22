package com.github9triver.cfn.config;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;

@Data
@ConfigurationProperties(prefix = "cfn.k8s")
public class K8sClientProperties {
    private String configFile;
    private String context;
}
