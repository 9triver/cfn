package com.github9triver.cfn;

import com.github9triver.cfn.config.K8sClientProperties;
import com.github9triver.cfn.config.NodeProperties;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;

@SpringBootApplication
@EnableConfigurationProperties({NodeProperties.class, K8sClientProperties.class})
public class CfnNodeApplication {

    public static void main(String[] args) {
        SpringApplication.run(CfnNodeApplication.class, args);
    }

}
