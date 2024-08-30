
#!bin/bash
{
sleep 1
echo -e "HSET\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\nname\nMaddux\n"

sleep 1
echo -e "HSET\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngroup\nMaddux's Group\n"

sleep 1
echo -e "HGET\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\nname\n"

sleep 1
echo -e "HGET\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngroup\n"

sleep 1
echo -e "HDEL\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\nname\n"

sleep 1
echo -e "HREM\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\n"

sleep 1
echo -e "RPUSH\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\nITALIAN\n"

sleep 1
echo -e "RPUSH\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\nAMERICAN\n"

sleep 1
echo -e "RPUSH\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\nJAPANESE\n"

sleep 1
echo -e "LRANGE\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\n0\n-1\n"

sleep 1
echo -e "LCLEAR\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\ngenres\n"

sleep 1
echo -e "SADD\nMaddux's Group\nMaddux\n"

sleep 1
echo -e "SADD\nMaddux's Group\nTrin\n"

sleep 1
echo -e "SGET\nMaddux's Group\n"

sleep 1
echo -e "SREM\nMaddux's Group\nTrin\n"

sleep 1
echo -e "SGET\nMaddux's Group\n"

sleep 1
echo -e "SREM\nMaddux's Group\nMaddux\n"

sleep 1
echo -e "SGET\nMaddux's Group\n"

sleep 1
echo -e "HREM\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\n"

sleep 1
echo -e "DEL\n843c1744-f6c2-6118-6a62-96ea50c2ea1d\n"
} | nc 127.0.0.1 6379
