
#!bin/bash
{
sleep 1
echo -e "HSET\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\nname\nMaddux\n"

sleep 1
echo -e "HSET\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngroup\nMaddux's Group\n"

sleep 1
echo -e "HSETLIST\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\nITALIAN, AMERICAN, JAPANESE\n"

sleep 1
echo -e "HGET\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\nname\n"

sleep 1
echo -e "HGET\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngroup\n"

sleep 1
echo -e "HGETLIST\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\n"

sleep 1
echo -e "HREMOVELISTFIELD\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\n"

sleep 1
echo -e "HREMOVESTRINGFIELD\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\nname\n"

sleep 1
echo -e "HREMOVE\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\n"

sleep 1
echo -e "RPUSH\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\nITALIAN\n"

sleep 1
echo -e "RPUSH\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\nAMERICAN\n"

sleep 1
echo -e "RPUSH\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\nJAPANESE\n"

sleep 1
echo -e "LRANGE\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\n"

} | nc 127.0.0.1 6379
