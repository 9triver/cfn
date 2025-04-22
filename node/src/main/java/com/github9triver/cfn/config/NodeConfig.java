package com.github9triver.cfn.config;

import com.github9triver.cfn.manager.LocalResourceManager;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.lang.reflect.InvocationTargetException;

@Configuration
public class NodeConfig {

    private final NodeProperties properties;

    public NodeConfig(NodeProperties properties) {
        this.properties = properties;
    }

    @Bean
    public LocalResourceManager localResourceManager() {
        LocalResourceManager manager;
        try {
            manager = properties.getLocalResourceManager().getConstructor().newInstance();
        } catch (InstantiationException | IllegalAccessException | InvocationTargetException |
                 NoSuchMethodException e) {
            throw new RuntimeException(e);
        }
        return manager;
    }


}
