package com.github9triver.cfn.manager;

import com.github9triver.cfn.model.dto.ResourceDto;
import com.github9triver.cfn.model.dto.ServerAddress;
import com.github9triver.cfn.proto.data.Resources;

public interface LocalResourceManager {

    ResourceDto getTotalResourceCount();

    ResourceDto getAvailableResourceCount();

    ServerAddress requestResources(Resources.Resource resource);
}
