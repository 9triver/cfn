package com.github9triver.cfn.util.mapper;

import com.github9triver.cfn.model.dto.ResourceDto;
import com.github9triver.cfn.model.vo.ResourceVo;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.LinkedHashMap;
import java.util.Map;

public class ResourceMapper {

    public static ResourceDto vo2dto(ResourceDto vo) {
        final ResourceDto dto = new ResourceDto();
        return dto;
    }

    public static ResourceVo dto2vo(final ResourceDto dto) {
        final ResourceVo vo = new ResourceVo();
        vo.getCpu().setCores(dto.getCpu().getCores().toPlainString());
        vo.getMemory().setCapacity(formatBinarySI(dto.getMemory().getCapacity()));
        return vo;
    }

    private static final BigDecimal Ki = BigDecimal.valueOf(1024);
    private static final BigDecimal Mi = Ki.multiply(Ki);
    private static final BigDecimal Gi = Mi.multiply(Ki);
    private static final BigDecimal Ti = Gi.multiply(Ki);

    private static String formatBinarySI(BigDecimal bytes) {
        Map<BigDecimal, String> units = new LinkedHashMap<>();
        units.put(Ti, "Ti");
        units.put(Gi, "Gi");
        units.put(Mi, "Mi");
        units.put(Ki, "Ki");

        for (Map.Entry<BigDecimal, String> entry : units.entrySet()) {
            if (bytes.compareTo(entry.getKey()) >= 0) {
                BigDecimal result = bytes.divide(entry.getKey(), 2, RoundingMode.HALF_UP);
                return result.stripTrailingZeros().toPlainString() + " " + entry.getValue();
            }
        }
        return bytes.stripTrailingZeros().toPlainString() + " bytes";
    }

}
