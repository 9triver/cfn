package com.github9triver.cfn.manager;

import com.github9triver.cfn.model.dto.ResourceDto;

public interface LocalResourceManager {

    ResourceDto getTotalResourceCount();

    ResourceDto getAvailableResourceCount();

}
