
import AspectRatio from '@mui/joy/AspectRatio';
import Card from '@mui/joy/Card';
import CardOverflow from '@mui/joy/CardOverflow';
import Divider from '@mui/joy/Divider';
import Typography from '@mui/joy/Typography';
import Link from '@mui/joy/Link';
import backgroundImg from '../../Images/background1.png'
import { Repository } from '../../Models/index'

export const RepoCard = (props: Repository) => {
    
    return (
      <div className="repoData">
        <Card variant="outlined" sx={{ width: 320, height: 290, maxHeight: '100%' }}>
          <CardOverflow>
            <AspectRatio ratio="2">
              <img
                src={backgroundImg}
                loading="lazy"
                alt=""
              />
            </AspectRatio>
          </CardOverflow>

          <CardOverflow sx={{display: 'flex', overflow: 'auto', overflowY: 'auto'}}>
            <CardOverflow>
              <Typography level="h2" sx={{ fontSize: 'md', mt: 2 }}>
                <Link href={props.Link.toString()} overlay underline="none">
                  {props.Name}
                </Link>
              </Typography>
              <Typography level="body2" sx={{ mt: 0.5, mb: 2 }}>
                {props.Language}
              </Typography>
            </CardOverflow>
            <Typography level="h2" sx={{ fontSize: 'md', mt: 2}}>
                {props.Description}
            </Typography>
          </CardOverflow>

          <Divider inset="context" />

          <CardOverflow
          variant="soft"
          sx={{
            display: 'flex',
            gap: 1.5,
            py: 1.5,
            px: 'var(--Card-padding)',
            bgcolor: 'rgba(255, 255, 255, 0.8)',
          }}
        >
          <Typography level="body3" sx={{ fontWeight: 'md', color: 'text.secondary' }}>
            Numer of ⭐'s {props.Stars.toString()}
          </Typography>
          <Divider orientation="vertical" />
          <Typography level="body3" sx={{ fontWeight: 'md', color: 'text.secondary' }}>
            {props.ReleaseDate.toString()}
          </Typography>

        </CardOverflow>
        </Card>
      </div>
      
    );
}